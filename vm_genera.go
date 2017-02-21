package gorgonia

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/chewxy/gorgonia/tensor"
	"github.com/pkg/errors"
)

type lispMachine struct {
	g *ExprGraph
	q []adInstr // a to-do list of differentiation instructions

	// state stuff, to allow continuation
	sorted Nodes
	fwd    int
	bwd    int

	// logging stuff
	watchlist Nodes
	logger    *log.Logger
	buf       *bytes.Buffer
	valueFmt  string
	tabcount  int
	logFlags  byte

	runFlags     byte // supposed to go into state stuff.  Placed here for better compacting of struct
	checkedRoots bool // supposed to go into state stuff.
}

func NewLispMachine(g *ExprGraph, opts ...VMOpt) *lispMachine {
	runFlags := (byte(0) | (byte(1) << fwdOnly)) | (1 << bwdOnly) // run fwd and backwards
	m := &lispMachine{
		g:        g,
		fwd:      -1,
		bwd:      -1,
		valueFmt: "%3.3f",
		logFlags: 0x0,      // log nothing
		runFlags: runFlags, // run only fwd and bwd
	}

	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *lispMachine) logBwd() bool { return (m.logFlags>>bwdOnly)&byte(1) == 1 }
func (m *lispMachine) doLogBwd()    { m.logFlags |= byte(1) << bwdOnly }
func (m *lispMachine) dontLogBwd()  { m.logFlags &= (^(byte(1) << bwdOnly)) }
func (m *lispMachine) runBwd() bool { return m.runFlags>>bwdOnly&byte(1) == 1 }
func (m *lispMachine) doExecBwd()   { m.runFlags |= byte(1) << bwdOnly }
func (m *lispMachine) dontExecBwd() { m.runFlags &= (^(byte(1) << bwdOnly)) }

func (m *lispMachine) logFwd() bool { return (m.logFlags>>fwdOnly)&byte(1) == 1 }
func (m *lispMachine) doLogFwd()    { m.logFlags |= byte(1) << fwdOnly }
func (m *lispMachine) dontLogFwd()  { m.logFlags &= (^(byte(1) << fwdOnly)) }
func (m *lispMachine) runFwd() bool { return m.runFlags>>fwdOnly&byte(1) == 1 }
func (m *lispMachine) doExecFwd()   { m.runFlags |= byte(1) << fwdOnly }
func (m *lispMachine) dontExecFwd() { m.runFlags &= (^(byte(1) << fwdOnly)) }

func (m *lispMachine) watchNaN() bool { return (m.runFlags>>watchNaN)&byte(1) == 1 }
func (m *lispMachine) doWatchNaN()    { m.runFlags |= byte(1) << watchNaN }
func (m *lispMachine) dontWatchNaN()  { m.runFlags &= (^(byte(1) << watchNaN)) }

func (m *lispMachine) watchInf() bool { return (m.runFlags>>watchInf)&byte(1) == 1 }
func (m *lispMachine) doWatchInf()    { m.runFlags |= byte(1) << watchInf }
func (m *lispMachine) dontWatchInf()  { m.runFlags &= (^(byte(1) << watchInf)) }

func (m *lispMachine) watchAll() bool { return (m.logFlags>>watchAll)&byte(1) == 1 }
func (m *lispMachine) doWatchAll()    { m.logFlags |= (byte(1) << watchAll) }
func (m *lispMachine) dontWatchAll()  { m.logFlags &= (^(byte(1) << watchAll)) }

func (m *lispMachine) dealloc() bool { return (m.runFlags>>allocVals)&byte(1) == 1 }
func (m *lispMachine) doDealloc()    { m.runFlags |= byte(1) << allocVals }
func (m *lispMachine) dontDealloc()  { m.runFlags &= (^(byte(1) << allocVals)) }

// check roots only applies if you want to run a backprop as well
func (m *lispMachine) checkRoots() (err error) {
	if !m.checkedRoots && m.runBwd() {
		machineLogf("Checking if provided graph is sensible")
		machineLogf("roots: %v", m.g.Roots())
		for _, root := range m.g.Roots() {
			if !root.IsScalar() && !root.isStmt {
				err = errors.Errorf("Expected cost to be a scalar. Got %v with shape %v instead", root, root.Shape())
				ioutil.WriteFile("err.dot", []byte(root.RestrictedToDot(2, 10)), 0644)
				return
			}
		}
	}
	return
}

func (m *lispMachine) prepGraph() (err error) {
	if m.sorted == nil {
		if m.sorted, err = Sort(m.g); err != nil {
			return errors.Wrap(err, sortFail)
		}

		m.fwd = len(m.sorted) - 1
	}
	return
}

func (m *lispMachine) forward() (err error) {
	if m.fwd < 0 {
		return nil // or err?
	}

	n := m.sorted[m.fwd]
	m.watchedLogf("n: %v (%x)", n, n.Hashcode())
	m.enterLoggingContext()
	defer m.leaveLoggingContext()

	if !n.isStmt {
		switch {
		case n.isArg():
			machineLogf("Unit() on input node")
			if err = n.bind(dvUnit(n.boundTo)); err != nil {
				return errors.Wrap(err, bindFail)
			}
			return
		case n.isRandom():
			machineLogf("binding value of random node")

			var v Value
			if v, err = n.op.Do(); err != nil {
				return errors.Wrapf(err, execFail, n.op, n)
			}

			// we wrap it in a dualValue, but don't allocate anything for the d
			if err = n.bind(dvUnit0(v)); err != nil {
				return errors.Wrap(err, bindFail)
			}
			return
		default:
			// do nothihng
		}
		m.watchedLogf(m.valueFmt, n.boundTo)
	}

	// other wise it's time to execute the op
	m.watchedLogf("execute Op")
	op := n.op
	var output *dualValue

	inputs := make([]*dualValue, len(n.children))
	for i, child := range n.children {
		dv := child.boundTo.(*dualValue)
		inputs[i] = dv
	}

	m.watchedLogf("Before:")
	m.watchedLogf(m.valueFmt, n.boundTo)

	switch {
	case (m.g.roots.Contains(n) || n.isRoot()) && !n.isStmt:
		machineLogf("Applying op %v to root", op)
		if n.boundTo == nil {
			machineLogf("dvBindVar")
			if output, err = dvBindVar(op, inputs); err != nil {
				return errors.Wrapf(err, execFail, op, n)
			}
			if err = n.bind(output); err != nil {
				return errors.Wrap(err, bindFail)
			}
		} else {
			machineLogf("dvBindVar0")
			dv, ok := n.boundTo.(*dualValue)
			if !ok {
				panic(fmt.Sprintf("n not dual value %v", n))
			}
			if err = dvBindVar0(op, dv, inputs); err != nil {
				return errors.Wrapf(err, execFail, op, n)
			}
		}

	case n.isStmt:
		machineLogf("ReadOp: %v ", op)
		switch ot := n.op.(type) {
		case readOp:
			childVal := n.children[0].boundTo
			if dv, ok := childVal.(*dualValue); ok {
				*ot.into = dv.Value
			} else {
				*ot.into = childVal
			}
		}

	case n.boundTo == nil:
		machineLogf("Fresh, unencountered node, so dvBind(%v)", op)
		machineLogf("Inputs")
		enterLoggingContext()
		for i, in := range inputs {
			if inT, ok := in.Value.(tensor.Tensor); ok {
				machineLogf("%d; %v", i, inT.Info())
			}
		}
		leaveLoggingContext()
		if output, err = dvBind(op, inputs); err != nil {
			return errors.Wrapf(err, execFail, op, n)
		}

		if err = n.bind(output); err != nil {
			return errors.Wrap(err, bindFail)
		}

	default:
		m.logf("bind(%v) with as much reuse as possible", op)
		// reuse as much as possible
		output := dvUnit(n.boundTo)
		if err = n.bind(output); err != nil {
			return errors.Wrap(err, bindFail)
		}

		err = dvBind0(op, output, inputs)
		if _, ok := errors.Cause(err).(AutoDiffError); ok {
			err = nil
		} else if err != nil {
			return errors.Wrapf(err, execFail, op, n)
		}
	}
	m.watchedLogf("After:")
	m.watchedLogf(m.valueFmt, n.boundTo)

	if aop, ok := op.(ADOp); ok && m.runBwd() {
		instr := adInstr{
			ADOp: aop,

			inputs: n.children,
			output: n,
		}
		m.q = append(m.q, instr)
	}

	if m.watchNaN() && !n.isStmt {
		if hasNaN(n.boundTo) {
			return errors.New("NaN found in value")
		}
	}

	return
}

func (m *lispMachine) backward() (err error) {
	if m.bwd < 0 {
		return errors.New("no backprop queue")
	}

	instr := m.q[m.bwd]
	m.watchedLogf("Differentiating op %v. Output: %v (%x)", instr, instr.output, instr.output.Hashcode())
	m.enterLoggingContext()
	defer m.leaveLoggingContext()

	m.watchedLogf("Inputs: %v", instr.inputs)
	m.enterLoggingContext()
	for _, in := range instr.inputs {
		m.watchedLogf(m.valueFmt, in.boundTo.(*dualValue).d)
	}
	m.leaveLoggingContext()

	// actual differentiation
	if err = instr.do(); err != nil {
		return errors.Wrapf(err, autodiffFail, instr.ADOp)
	}

	m.watchedLogf("After:")
	m.enterLoggingContext()
	for _, in := range instr.inputs {
		m.watchedLogf(m.valueFmt, in.boundTo.(*dualValue).d)
	}

	m.leaveLoggingContext()

	if m.watchNaN() {
		if hasNaN(instr.output.boundTo) {
			return errors.New("NaN found in value")
		}

		for _, in := range instr.inputs {
			if hasNaN(in.boundTo) {
				return errors.New("NaN found in value")
			}
		}
	}
	return
}

func (m *lispMachine) RunAll() (err error) {
	if err = m.prepGraph(); err != nil {
		return errors.Wrap(err, "Could not prepGraph()")
	}

	if err = m.checkRoots(); err != nil {
		return errors.Wrap(err, "Could not checkRoots()")
	}

	if !m.runFwd() {
		goto backward
	}

	for err = nil; err == nil && m.fwd >= 0; m.fwd-- {
		err = m.forward()
	}

	if err != nil {
		return
	}

backward:
	if !m.runBwd() {
		return nil
	}

	if m.bwd < 0 {
		m.bwd = len(m.q) - 1
	}

	for err = nil; err == nil && m.bwd >= 0; m.bwd-- {
		err = m.backward()
	}
	return
}

func (m *lispMachine) Free() {
	if m.dealloc() {
		for _, n := range m.sorted {
			m.logf("dealloc n; %v %x %p", n, n.Hashcode(), n)
			if !n.isInput() {
				n.unbind()
				m.logf("OK")
			}
		}
	}
}

func (m *lispMachine) Reset() {
	m.fwd = len(m.sorted) - 1
	m.bwd = len(m.q) - 1
}

func (m *lispMachine) LastRun() (n *Node, backprop bool) {
	if m.fwd < 0 && m.runBwd() {
		goto backward
	} else if !m.runBwd() {
		n = m.sorted[0] // last to run
		return
	} else {
		n = m.sorted[m.fwd]
		return
	}

backward:
	backprop = true
	if m.bwd < 0 {
		n = m.q[0].output
		return
	}
	n = m.q[m.bwd].output
	return
}

func (m *lispMachine) watchedLogf(format string, attrs ...interface{}) {
	if !m.logFwd() && !DEBUG {
		goto backwards
	}

	if m.fwd >= 0 {
		n := m.sorted[m.fwd]
		if m.watchlist.Contains(n) || m.watchAll() || DEBUG {
			m.logf(format, attrs...)
		}
		return
	}

backwards:
	if !m.logBwd() && !DEBUG {
		return
	}

	if m.bwd >= 0 {
		instr := m.q[m.bwd]
		write := m.watchlist.Contains(instr.output)
		if !write {
			for _, in := range instr.inputs {
				if m.watchlist.Contains(in) {
					write = true
					break
				}
			}
		}

		if write || m.watchAll() || DEBUG {
			m.logf(format, attrs...)
		}
	}
}

func (m *lispMachine) logf(format string, attrs ...interface{}) {
	switch {
	case machineDev, autodiffDev:
		if machineDev {

			machineLogf(format, attrs...)
		} else {
			autodiffLogf(format, attrs...)
		}

		if m.logger != nil {
			goto loggercase
		}

		break

	loggercase:
		fallthrough
	case m.logger != nil:
		s := fmt.Sprintf(format, attrs...)
		s = strings.Replace(s, "\n", m.buf.String(), -1)
		m.logger.Println(s)
	}
}

func (m *lispMachine) enterLoggingContext() {
	if DEBUG && machineDev {
		enterLoggingContext()
	}
	m.tabcount++
	if m.logger != nil {
		reps := strings.Repeat("\t", m.tabcount)
		m.logger.SetPrefix(reps)
		m.buf.Reset()
		m.buf.WriteString("\n")
		m.buf.WriteString(reps)
	}
}

func (m *lispMachine) leaveLoggingContext() {
	if DEBUG && machineDev {
		leaveLoggingContext()
	}
	m.tabcount--
	if m.tabcount < 0 {
		m.tabcount = 0
	}
	if m.logger != nil {
		reps := strings.Repeat("\t", m.tabcount)
		m.logger.SetPrefix(reps)
		m.buf.Reset()
		m.buf.WriteString("\n")
		m.buf.WriteString(reps)
	}
}

// adInstr is an autodifferentiation instruction
type adInstr struct {
	ADOp

	inputs Nodes
	output *Node
}

func (instr adInstr) do() error {
	return instr.ADOp.DoDiff(instr.inputs, instr.output)
}
