package gorgonia

import (
	"math"

	tf32 "github.com/chewxy/gorgonia/tensor/f32"
	tf64 "github.com/chewxy/gorgonia/tensor/f64"
	"github.com/chewxy/gorgonia/tensor/types"
	"github.com/chewxy/math32"
)

// Steps is anything that does gradient updates.
// The name solvers is stolen from Caffe. A much shorter name than GradientUpdaters
type Steps interface {
	Step(Nodes) error
}

// Solver is a struct that lets you create a Solver with l1reg, l2reg, clip & Steps
// To create a 'instance' of solver It's needed to do something like this:
// Solver{
//   Steps: &Steps,
//   l1reg: ...,
//   ...
// }
type Solver struct {
	Steps Steps
	l1reg float64 // l2 regularization
	l2reg float64 // l2 regularization
	clip  float64 // clip value
}

// SolverOpt is a function that provides construction options for a Solver
type SolverOpt func(s Solver)

// WithL2Reg adds a L2 regularization parameter to the solver. By default, the solvers do not use any regularization param
func WithL2Reg(l2reg float64) SolverOpt {
	return func(s Solver) {
		switch st := s.(type) {
		case *RMSPropSolver:
			st.l2reg = l2reg
			st.useL2Reg = true
		case *AdamSolver:
			st.l2reg = l2reg
			st.useL2Reg = true
		case *VanillaSolver:
			st.l2reg = l2reg
			st.useL2Reg = true
		}
	}
}

// WithL1Reg adds a L1 regularization parameter to the solver. By default, the solvers do not use any regularization param
func WithL1Reg(l1reg float64) SolverOpt {
	return func(s Solver) {
		switch st := s.(type) {
		case *AdamSolver:
			st.l1reg = l1reg
			st.useL1Reg = true
		case *VanillaSolver:
			st.l1reg = l1reg
			st.useL1Reg = true
		}
	}
}

// WithBatchSize sets the batch size for the solver. Currently only Adam and Vanilla (basic SGD) has batch size support
func WithBatchSize(batch float64) SolverOpt {
	return func(s Solver) {
		switch st := s.(type) {
		case *AdamSolver:
			st.batch = batch
		case *VanillaSolver:
			st.batch = batch
		}
	}
}

// WithEps sets the smoothing factor for the solver.
func WithEps(eps float64) SolverOpt {
	return func(s Solver) {
		switch st := s.(type) {
		case *RMSPropSolver:
			st.eps = eps
		case *AdamSolver:
			st.eps = eps
		}
	}
}

// WithClip clips the gradient if it gets too crazy. By default all solvers do not have any clips attached
func WithClip(clip float64) SolverOpt {
	return func(s Solver) {
		switch st := s.(type) {
		case *RMSPropSolver:
			st.clip = clip
			st.useClip = true
		case *AdamSolver:
			st.clip = clip
			st.useClip = true
		case *VanillaSolver:
			st.clip = clip
			st.useClip = true
		}
	}
}

// WithLearnRate sets the learn rate or step size for the solver.
func WithLearnRate(eta float64) SolverOpt {
	return func(s Solver) {
		switch st := s.(type) {
		case *RMSPropSolver:
			st.eta = eta
		case *AdamSolver:
			st.eta = eta
		case *VanillaSolver:
			st.eta = eta
		}
	}
}

// WithBeta1 sets the beta1 param of the solver. Only works with Adam
func WithBeta1(beta1 float64) SolverOpt {
	return func(s Solver) {
		switch st := s.(type) {
		case *AdamSolver:
			st.beta1 = beta1
		}
	}
}

// WithBeta2 sets the beta1 param of the solver. Only works with Adam
func WithBeta2(beta2 float64) SolverOpt {
	return func(s Solver) {
		switch st := s.(type) {
		case *AdamSolver:
			st.beta2 = beta2
		}
	}
}

// WithRho sets the decay parameter of the RMSProp solver
func WithRho(rho float64) SolverOpt {
	return func(s Solver) {
		switch st := s.(type) {
		case *RMSPropSolver:
			st.decay = rho
		}
	}
}

// http://www.cs.toronto.edu/~tijmen/csc321/slides/lecture_slides_lec6.pdf
type RMSPropSolver struct {
	decay float64 // decay rate/rho
	eps   float64 // smoothing factor
	l2reg float64 // l2 regularization
	clip  float64 // clip value
	eta   float64 // learn rate

	useClip, useL2Reg bool

	// unsettable
	cache []*dualValue
}

// NewRMSPropSolver creates an RMSProp solver with these default values:
//		eta (learn rate)	  : 0.001
//		eps (smoothing factor): 1e-8
//		rho (decay factor)    : 0.999
func NewRMSPropSolver(opts ...SolverOpt) *RMSPropSolver {
	s := &RMSPropSolver{
		decay: 0.999,
		eps:   1e-8,
		eta:   0.001,
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *RMSPropSolver) Step(model Nodes) (err error) {
	if s.cache == nil {
		s.cache = make([]*dualValue, len(model))
	}

	for i, n := range model {
		solverLogf("BEFORE (%v %p) : %+1.1s", n, n, n.boundTo)
		dv, ok := n.boundTo.(*dualValue)
		if !ok {
			err = NewError(typeError, "Expected a *dualValue in %v (%x). Got %T instead", n, n.Hashcode(), n.boundTo)
			return
		}

		var cached *dualValue
		if cached = s.cache[i]; cached == nil {
			if cached, err = dv.clone0(); err != nil {
				return
			}
			s.cache[i] = cached
		}

		var dt Dtype
		dt, err = dtypeOf(dv.Type())
		if err != nil {
			return
		}

		grad := dv.d
		weights := dv.Value

		cv := cached.Value
		// cw = cw*decay + (1-decay) * grad^2
		switch cw := cv.(type) {
		case Tensor:
			switch dt {
			case Float32:
				var gt, gt2, w, regularized *tf32.Tensor
				decay := float32(s.decay)
				omdecay := float32(1.0 - s.decay)
				stepSize := float32(s.eta)
				eps := float32(s.eps)
				l2reg := float32(s.l2reg)
				clip := float32(s.clip)

				gt = grad.(Tensor).Tensor.(*tf32.Tensor)
				if gt2, err = tf32.PointwiseSquare(gt); err != nil { // safe version
					return
				}

				t := cw.Tensor.(*tf32.Tensor)
				tf32.PointwiseMul(t, decay, types.UseUnsafe())
				tf32.PointwiseMul(gt2, omdecay, types.UseUnsafe())
				tf32.Add(t, gt2, types.UseUnsafe())

				if s.useClip {
					if _, err = tf32.Clamp(gt, -clip, clip, types.UseUnsafe()); err != nil {
						return
					}
				}

				// update and regularize
				var upd *tf32.Tensor
				if upd, err = tf32.Add(t, eps); err != nil {
					return
				}

				if _, err = tf32.InvSqrt(upd, types.UseUnsafe()); err != nil {
					return
				}
				if _, err = tf32.PointwiseMul(gt, -stepSize, types.UseUnsafe()); err != nil {
					return
				}

				if _, err = tf32.PointwiseMul(upd, gt, types.UseUnsafe()); err != nil {
					return
				}

				w = weights.(Tensor).Tensor.(*tf32.Tensor)

				if s.useL2Reg {
					if regularized, err = tf32.PointwiseMul(w, l2reg); err != nil {
						return
					}

					if _, err = tf32.Sub(upd, regularized, types.UseUnsafe()); err != nil {
						return
					}
				}

				if _, err = tf32.Add(w, upd, types.UseUnsafe()); err != nil {
					return
				}

				// zero all
				gt.Zero()

			case Float64:
				var gt, gt2, w, regularized *tf64.Tensor
				decay := s.decay
				omdecay := 1.0 - s.decay
				stepSize := s.eta
				eps := s.eps
				l2reg := s.l2reg
				clip := s.clip

				gt = grad.(Tensor).Tensor.(*tf64.Tensor)
				if gt2, err = tf64.PointwiseSquare(gt); err != nil { // safe version
					return
				}

				t := cw.Tensor.(*tf64.Tensor)
				tf64.PointwiseMul(t, decay, types.UseUnsafe())
				tf64.PointwiseMul(gt2, omdecay, types.UseUnsafe())
				tf64.Add(t, gt2, types.UseUnsafe())

				if s.useClip {
					if _, err = tf64.Clamp(gt, -clip, clip, types.UseUnsafe()); err != nil {
						return
					}
				}

				// update and regularize
				var upd *tf64.Tensor
				if upd, err = tf64.Add(t, eps); err != nil {
					return
				}

				if _, err = tf64.InvSqrt(upd, types.UseUnsafe()); err != nil {
					return
				}
				if _, err = tf64.PointwiseMul(gt, -stepSize, types.UseUnsafe()); err != nil {
					return
				}

				if _, err = tf64.PointwiseMul(upd, gt, types.UseUnsafe()); err != nil {
					return
				}

				w = weights.(Tensor).Tensor.(*tf64.Tensor)

				if s.useL2Reg {
					if regularized, err = tf64.PointwiseMul(w, l2reg); err != nil {
						return
					}

					if _, err = tf64.Sub(upd, regularized, types.UseUnsafe()); err != nil {
						return
					}
				}

				if _, err = tf64.Add(w, upd, types.UseUnsafe()); err != nil {
					return
				}

				// zero all
				gt.Zero()

			default:
				err = NewError(NotYetImplemented, "Tensor of %T not implemented yet for RMSProp", dt)
				return
			}

		case Scalar:
			switch dt {
			case Float32:
				decay := float32(s.decay)
				omdecay := float32(1.0 - s.decay)
				stepSize := float32(s.eta)
				eps := float32(s.eps)
				l2reg := float32(s.l2reg)

				gs := grad.(Scalar).v.(float32)
				c := cw.v.(float32)
				c = c*decay + omdecay*gs*gs

				w := weights.(Scalar).v.(float32)
				upd := -stepSize*gs/float32(math.Sqrt(float64(c+eps))) - l2reg*w
				w += upd

				// because scalar values are copies, and not pointers, we have to actually re-update the dualValu in model[i]
				var ws Value
				if ws, err = anyToValue(w); err != nil {
					return
				}

				dv.Value = ws
				dv.d = Scalar{t: Float32, v: float32(0)} // zero it
			case Float64:
				decay := s.decay
				omdecay := 1.0 - s.decay
				stepSize := s.eta
				eps := s.eps
				l2reg := s.l2reg

				gs := grad.(Scalar).v.(float64)
				c := cw.v.(float64)
				c = c*decay + omdecay*gs*gs

				w := weights.(Scalar).v.(float64)
				upd := -stepSize*gs/math.Sqrt(c+eps) - l2reg*w
				w += upd

				// because scalar values are copies, and not pointers, we have to actually re-update the dualValu in model[i]
				var ws Value
				if ws, err = anyToValue(w); err != nil {
					return
				}

				dv.Value = ws
				dv.d = Scalar{t: Float64, v: float64(0)} // zero it
			default:
				err = NewError(NotYetImplemented, "Scalar of %T not yet implemented for RMSProp", dt)
				return
			}
		}
		solverLogf("AFTER (%v): %+1.1s", n, n.boundTo)
	}
	return nil
}

// AdamSolver is the Adaptive Moment Estimation solver (basically RMSProp on steroids).
// Paper: http://arxiv.org/abs/1412.6980
//
// We overload the purpose of existing data structure of a *dualValue. However, instead of just holding a value and its derivative,
// the cache's *dualValues hold the Means of gradients (in .Value) and the variances of the gradients (in .d)
type AdamSolver struct {
	eta   float64 // learn rate
	eps   float64 // smoothing
	beta1 float64 // modifier for means
	beta2 float64 // modifier for variances
	clip  float64 // clip gradients
	l1reg float64 // l1 regularization parameter
	l2reg float64 // l2 regularization parameter
	batch float64 // batch size

	useClip, useL1Reg, useL2Reg bool

	// unsettable
	iter  int
	cache []*dualValue
}

// NewAdamSolver creates an Adam solver with these default values:
//		eta (learn rate)	  	: 0.001
//		eps (smoothing factor)		: 1e-8
//		beta1				: 0.9
//		beta2 				: 0.999
//		batch				: 1
func NewAdamSolver(opts ...SolverOpt) *AdamSolver {
	s := &AdamSolver{
		eta:   0.001,
		eps:   1e-8,
		beta1: 0.9,
		beta2: 0.999,
		batch: 1,
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *AdamSolver) Step(model Nodes) (err error) {
	if s.cache == nil {
		s.cache = make([]*dualValue, len(model))
	}

	s.iter++
	correction1 := (1 - math.Pow(s.beta1, float64(s.iter)))
	correction2 := (1 - math.Pow(s.beta2, float64(s.iter)))

	for i, n := range model {
		dv, ok := n.boundTo.(*dualValue)
		if !ok {
			err = NewError(typeError, "Expected a *dualValue in %v (%x). Got %T instead", n, n.Hashcode(), n.boundTo)
			return
		}

		var cached *dualValue
		if cached = s.cache[i]; cached == nil {
			if cached, err = dv.clone0(); err != nil {
				return
			}
			s.cache[i] = cached
		}

		var dt Dtype
		if dt, err = dtypeOf(dv.Type()); err != nil {
			return
		}

		grad := dv.d
		weights := dv.Value

		cvm := cached.Value // means of gradients
		cvv := cached.d     // variances of gradients

		switch cm := cvm.(type) {
		case Tensor:
			switch dt {
			case Float32:
				cv := cvv.(Tensor)
				g := grad.(Tensor).Tensor.(*tf32.Tensor)
				w := weights.(Tensor).Tensor.(*tf32.Tensor)
				m := cm.Tensor.(*tf32.Tensor)
				v := cv.Tensor.(*tf32.Tensor)

				l1reg := float32(s.l1reg)
				l2reg := float32(s.l2reg)
				batch := float32(s.batch)
				clip := float32(s.clip)
				beta1 := float32(s.beta1)
				beta2 := float32(s.beta2)
				eps := float32(s.eps)
				eta := float32(s.eta)

				// prep the regularization of gradients
				var l1regs, l2regs *tf32.Tensor
				if s.useL1Reg {
					if l1regs, err = tf32.Sign(w); err != nil {
						return
					}

					if l1regs, err = tf32.PointwiseMul(l1reg, l1regs, types.UseUnsafe()); err != nil {
						return
					}

					if g, err = tf32.Add(g, l1regs, types.UseUnsafe()); err != nil {
						return
					}

					defer tf32.ReturnTensor(l1regs)
				}

				if s.useL2Reg {
					if l2regs, err = tf32.PointwiseMul(l2reg, w); err != nil {
						return
					}

					if g, err = tf32.Add(g, l2regs, types.UseUnsafe()); err != nil {
						return
					}

					defer tf32.ReturnTensor(l2regs)
				}

				if batch > 1 {
					if g, err = tf32.PointwiseMul(1/batch, g, types.UseUnsafe()); err != nil {
						return
					}
				}

				if s.useClip && clip > 0 {
					if g, err = tf32.Clamp(g, -clip, clip, types.UseUnsafe()); err != nil {
						return
					}
				}

				// prep done. Now let's apply the formula:
				// the formula is
				//		(β_1 * m_t-1) + (1 - β_1)g_t ..................	1
				//		(β_2 * v_t-1) + (1 - β_2)*(g_t)^2 .............	2

				// equation(1)
				t1 := g.Clone()
				if t1, err = tf32.PointwiseMul((1 - beta1), t1, types.UseUnsafe()); err != nil {
					return
				}

				// equation(2)
				if g, err = tf32.PointwiseMul(g, g, types.UseUnsafe()); err != nil {
					return
				}
				if g, err = tf32.PointwiseMul((1 - beta2), g, types.UseUnsafe()); err != nil {
					return
				}

				// equation (1)
				if t1, err = tf32.PointwiseMul(beta1, m, types.WithIncr(t1)); err != nil {
					return
				}

				// equation (2)
				if g, err = tf32.PointwiseMul(beta2, v, types.WithIncr(g)); err != nil {
					return
				}

				cm.Tensor = t1
				cv.Tensor = g.Clone() // g belongs to the node's dual value, so clone here
				defer tf32.ReturnTensor(m)
				defer tf32.ReturnTensor(v)

				// now deal with the hats
				mHats := t1.Clone()
				vHats := g.Clone()

				if mHats, err = tf32.PointwiseMul((float32(1) / float32(correction1)), mHats, types.UseUnsafe()); err != nil {
					return
				}

				if vHats, err = tf32.PointwiseMul((float32(1) / float32(correction2)), vHats, types.UseUnsafe()); err != nil {
					return
				}

				// update := -eta * mHat / (sqrt(vHat) + epsilon)
				if vHats, err = tf32.Sqrt(vHats, types.UseUnsafe()); err != nil {
					return // TODO: rewrite this to use InvSqrt
				}

				if vHats, err = tf32.Add(eps, vHats, types.UseUnsafe()); err != nil {
					return
				}

				if mHats, err = tf32.PointwiseMul(-eta, mHats, types.UseUnsafe()); err != nil {
					return
				}

				if w, err = tf32.PointwiseDiv(mHats, vHats, types.WithIncr(w)); err != nil {
					return
				}

				defer tf32.ReturnTensor(vHats)
				defer tf32.ReturnTensor(mHats)

				if _, err = tf64.Add(w, mHats, types.UseUnsafe()); err != nil {
					return
				}

				g.Zero()
			case Float64:
				cv := cvv.(Tensor)
				g := grad.(Tensor).Tensor.(*tf64.Tensor)
				w := weights.(Tensor).Tensor.(*tf64.Tensor)
				m := cm.Tensor.(*tf64.Tensor)
				v := cv.Tensor.(*tf64.Tensor)

				l1reg := s.l1reg
				l2reg := s.l2reg
				batch := s.batch
				clip := s.clip
				beta1 := s.beta1
				beta2 := s.beta2
				eps := s.eps
				eta := s.eta

				// prep the regularization of gradients
				var l1regs, l2regs *tf64.Tensor
				if s.useL1Reg {
					if l1regs, err = tf64.Sign(w); err != nil {
						return
					}

					if l1regs, err = tf64.PointwiseMul(l1reg, l1regs, types.UseUnsafe()); err != nil {
						return
					}

					if g, err = tf64.Add(g, l1regs, types.UseUnsafe()); err != nil {
						return
					}

					defer tf64.ReturnTensor(l1regs)
				}

				if s.useL2Reg {
					if l2regs, err = tf64.PointwiseMul(l2reg, w); err != nil {
						return
					}

					if g, err = tf64.Add(g, l2regs, types.UseUnsafe()); err != nil {
						return
					}

					defer tf64.ReturnTensor(l2regs)
				}

				if batch > 1 {
					if g, err = tf64.PointwiseMul(1/batch, g, types.UseUnsafe()); err != nil {
						return
					}
				}

				if s.useClip && clip > 0 {
					if g, err = tf64.Clamp(g, -clip, clip, types.UseUnsafe()); err != nil {
						return
					}
				}

				// prep done. Now let's apply the formula:
				// the formula is
				//		(β_1 * m_t-1) + (1 - β_1)g_t ..................	1
				//		(β_2 * v_t-1) + (1 - β_2)*(g_t)^2 .............	2

				// equation(1)
				t1 := g.Clone()
				if t1, err = tf64.PointwiseMul((1 - beta1), t1, types.UseUnsafe()); err != nil {
					return
				}

				// equation(2)
				if g, err = tf64.PointwiseMul(g, g, types.UseUnsafe()); err != nil {
					return
				}
				if g, err = tf64.PointwiseMul((1 - beta2), g, types.UseUnsafe()); err != nil {
					return
				}

				// equation (1)
				if t1, err = tf64.PointwiseMul(beta1, m, types.WithIncr(t1)); err != nil {
					return
				}

				// equation (2)
				if g, err = tf64.PointwiseMul(beta2, v, types.WithIncr(g)); err != nil {
					return
				}

				cm.Tensor = t1
				cv.Tensor = g.Clone() // g belongs to the node's dual value, so clone here
				defer tf64.ReturnTensor(m)
				defer tf64.ReturnTensor(v)

				// now deal with the hats
				mHats := t1.Clone()
				vHats := g.Clone()

				if mHats, err = tf64.PointwiseMul((1 / correction1), mHats, types.UseUnsafe()); err != nil {
					return
				}

				if vHats, err = tf64.PointwiseMul((1 / correction2), vHats, types.UseUnsafe()); err != nil {
					return
				}

				// update := -eta * mHat / (sqrt(vHat) + epsilon)
				if vHats, err = tf64.Sqrt(vHats, types.UseUnsafe()); err != nil {
					return // TODO: rewrite this to use InvSqrt
				}

				if vHats, err = tf64.Add(eps, vHats, types.UseUnsafe()); err != nil {
					return
				}

				if mHats, err = tf64.PointwiseMul(-eta, mHats, types.UseUnsafe()); err != nil {
					return
				}

				if w, err = tf64.PointwiseDiv(mHats, vHats, types.WithIncr(w)); err != nil {
					return
				}

				defer tf64.ReturnTensor(vHats)
				defer tf64.ReturnTensor(mHats)

				if _, err = tf64.Add(w, mHats, types.UseUnsafe()); err != nil {
					return
				}

				g.Zero()
			default:
				err = NewError(NotYetImplemented, "Tensor of %T not implemented yet for AdamSolver", dt)
				return
			}
		case Scalar:
			switch dt {
			case Float32:
				cv := cvv.(Scalar)
				g := grad.(Scalar).v.(float32)
				w := weights.(Scalar).v.(float32)
				m := cm.v.(float32)
				v := cv.v.(float32)

				l1reg := float32(s.l1reg)
				l2reg := float32(s.l2reg)
				batch := float32(s.batch)
				clip := float32(s.clip)
				beta1 := float32(s.beta1)
				beta2 := float32(s.beta2)
				eps := float32(s.eps)
				eta := float32(s.eta)

				if s.useL1Reg {
					if w < 0 {
						l1reg = -l1reg
					}
					g += l1reg
				}

				if s.useL2Reg {
					l2reg *= w
					g += l2reg
				}

				if batch > 1 {
					g *= (1 / batch)
				}

				if s.useClip {
					if g > clip {
						g = clip
					} else if g < -clip {
						g = -clip
					}
				}

				newM := (beta1 * m) + (1-beta1)*g
				newV := (beta2 * v) + (1-beta2)*g*g

				cv.v = newV
				cm.v = newM

				mHat := (1 / float32(correction1)) * newM
				vHat := (1 / float32(correction2)) * newV

				upd := -eta * mHat / (float32(math.Sqrt(float64(vHat))) + eps)
				w += upd

				dv.Value = Scalar{t: Float32, v: w}
				dv.d = Scalar{t: Float32, v: float32(0.0)}
			case Float64:
				cv := cvv.(Scalar)
				g := grad.(Scalar).v.(float64)
				w := weights.(Scalar).v.(float64)
				m := cm.v.(float64)
				v := cv.v.(float64)

				l1reg := s.l1reg
				l2reg := s.l2reg
				batch := s.batch
				clip := s.clip
				beta1 := s.beta1
				beta2 := s.beta2
				eps := s.eps
				eta := s.eta

				if s.useL1Reg {
					if w < 0 {
						l1reg = -l1reg
					}
					g += l1reg
				}

				if s.useL2Reg {
					l2reg *= w
					g += l2reg
				}

				if batch > 1 {
					g *= (1 / batch)
				}

				if s.useClip {
					if g > clip {
						g = clip
					} else if g < -clip {
						g = -clip
					}
				}

				newM := (beta1 * m) + (1-beta1)*g
				newV := (beta2 * v) + (1-beta2)*g*g

				cv.v = newV
				cm.v = newM

				mHat := (1 / correction1) * newM
				vHat := (1 / correction2) * newV

				upd := -eta * mHat / (math.Sqrt(vHat) + eps)
				w += upd

				dv.Value = Scalar{t: Float64, v: w}
				dv.d = Scalar{t: Float64, v: 0.0}

			default:
				err = NewError(NotYetImplemented, "Scalar of %T not implemented yet for AdamSolver", dt)
				return
			}
		}

	}
	return
}

// VanillaSolver is your bog standard stochastic gradient descent optimizer. There are no fancy features to this
type VanillaSolver struct {
	eta   float64 // learn rate
	clip  float64 // clip gradients
	l1reg float64 // l1 regularization parameter
	l2reg float64 // l2 regularization parameter
	batch float64 // batch size

	useClip, useL1Reg, useL2Reg bool
}

// NewVanillaSolver creates a new VanillaSolver with sane-ish default values
func NewVanillaSolver(opts ...SolverOpt) *VanillaSolver {
	s := &VanillaSolver{
		batch: 1,
		eta:   0.001,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *VanillaSolver) Step(model Nodes) (err error) {
	for _, n := range model {
		dv, ok := n.boundTo.(*dualValue)
		if !ok {
			err = NewError(typeError, "Expected a *dualValue in %v (%x). Got %T instead", n, n.Hashcode(), n.boundTo)
			return
		}

		var dt Dtype
		dt, err = dtypeOf(dv.Type())
		if err != nil {
			return
		}

		grad := dv.d
		weights := dv.Value

		switch wt := weights.(type) {
		case Tensor:
			switch dt {
			case Float32:
				w := wt.Tensor.(*tf32.Tensor)
				g := grad.(Tensor).Tensor.(*tf32.Tensor)

				l1reg := float32(s.l1reg)
				l2reg := float32(s.l2reg)
				batch := float32(s.batch)
				clip := float32(s.clip)
				eta := float32(s.eta)

				// prep the regularization of gradients
				var l1regs, l2regs *tf32.Tensor
				if s.useL1Reg {
					if l1regs, err = tf32.Sign(w); err != nil {
						return
					}

					if l1regs, err = tf32.PointwiseMul(l1reg, l1regs, types.UseUnsafe()); err != nil {
						return
					}

					if g, err = tf32.Add(g, l1regs, types.UseUnsafe()); err != nil {
						return
					}

					defer tf32.ReturnTensor(l1regs)
				}

				if s.useL2Reg {
					if l2regs, err = tf32.PointwiseMul(l2reg, w); err != nil {
						return
					}

					if g, err = tf32.Add(g, l2regs, types.UseUnsafe()); err != nil {
						return
					}

					defer tf32.ReturnTensor(l2regs)
				}

				if batch > 1 {
					if g, err = tf32.PointwiseMul(1/batch, g, types.UseUnsafe()); err != nil {
						return
					}
				}

				if s.useClip && clip > 0 {
					if g, err = tf32.Clamp(g, -clip, clip, types.UseUnsafe()); err != nil {
						return
					}
				}

				if g, err = tf32.PointwiseMul(-eta, g, types.UseUnsafe()); err != nil {
					return
				}

				if _, err = tf32.Add(w, g, types.UseUnsafe()); err != nil {
					return
				}

				g.Zero()
			case Float64:
				w := wt.Tensor.(*tf64.Tensor)
				g := grad.(Tensor).Tensor.(*tf64.Tensor)

				l1reg := s.l1reg
				l2reg := s.l2reg
				batch := s.batch
				clip := s.clip
				eta := s.eta

				// prep the regularization of gradients
				var l1regs, l2regs *tf64.Tensor
				if s.useL1Reg {
					if l1regs, err = tf64.Sign(w); err != nil {
						return
					}

					if l1regs, err = tf64.PointwiseMul(l1reg, l1regs, types.UseUnsafe()); err != nil {
						return
					}

					if g, err = tf64.Add(g, l1regs, types.UseUnsafe()); err != nil {
						return
					}

					defer tf64.ReturnTensor(l1regs)
				}

				if s.useL2Reg {
					if l2regs, err = tf64.PointwiseMul(l2reg, w); err != nil {
						return
					}

					if g, err = tf64.Add(g, l2regs, types.UseUnsafe()); err != nil {
						return
					}

					defer tf64.ReturnTensor(l2regs)
				}

				if batch > 1 {
					if g, err = tf64.PointwiseMul(1/batch, g, types.UseUnsafe()); err != nil {
						return
					}
				}

				if s.useClip && clip > 0 {
					if g, err = tf64.Clamp(g, -clip, clip, types.UseUnsafe()); err != nil {
						return
					}
				}

				if g, err = tf64.PointwiseMul(-eta, g, types.UseUnsafe()); err != nil {
					return
				}

				if _, err = tf64.Add(w, g, types.UseUnsafe()); err != nil {
					return
				}

				g.Zero()

			default:
				err = nyi("VanillaSolver.Step", wt)
				return
			}
		case Scalar:
			switch dt {
			case Float32:
				g := grad.(Scalar).v.(float32)
				w := wt.v.(float32)

				l1reg := float32(s.l1reg)
				l2reg := float32(s.l2reg)
				batch := float32(s.batch)
				clip := float32(s.clip)
				eta := float32(s.eta)

				if s.useL1Reg {
					if w < 0 {
						l1reg = -l1reg
					}
					g += l1reg
				}

				if s.useL2Reg {
					l2reg *= w
					g += l2reg
				}

				if batch > 1 {
					g *= (1 / batch)
				}

				if s.useClip {
					if g > clip {
						g = clip
					} else if g < -clip {
						g = -clip
					}
				}

				upd := -eta * g
				w += upd

				dv.Value = Scalar{t: Float32, v: w}
				dv.d = Scalar{t: Float32, v: float32(0.0)}
			case Float64:
				g := grad.(Scalar).v.(float64)
				w := wt.v.(float64)

				l1reg := s.l1reg
				l2reg := s.l2reg
				batch := s.batch
				clip := s.clip
				eta := s.eta

				if s.useL1Reg {
					if w < 0 {
						l1reg = -l1reg
					}
					g += l1reg
				}

				if s.useL2Reg {
					l2reg *= w
					g += l2reg
				}

				if batch > 1 {
					g *= (1 / batch)
				}

				if s.useClip {
					if g > clip {
						g = clip
					} else if g < -clip {
						g = -clip
					}
				}

				upd := -eta * g
				w += upd

				dv.Value = Scalar{t: Float64, v: w}
				dv.d = Scalar{t: Float64, v: 0.0}
			default:
				err = nyi("VanillaSolver.step", wt)
				return
			}
		default:
			err = nyi("VanillaSolver.step", wt)
			return
		}
	}
	return
}

// AdaGradSolver is the solver that does adaptive gradient descent. Read the paper: http://jmlr.org/papers/v12/duchi11a.html
type AdaGradSolver struct {
	eta   float64 // learn rate
	eps   float64 // smoothing factor
	l1Reg float64 // l1reg param
	l2reg float64 // l2reg param
	clip  float64 // clip at

	useL2Reg, useClip bool

	cache []*dualValue
}

// NewAdaGradSolver creates a new AdaGradSolver with sane-ish default values
func NewAdaGradSolver(opts ...SolverOpt) *AdaGradSolver {
	s := &AdaGradSolver{
		eta: 0.001,
		eps: 1e-8,
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *AdaGradSolver) Step(model Nodes) (err error) {
	if s.cache == nil {
		s.cache = make([]*dualValue, len(model))
	}

	for i, n := range model {
		dv, ok := n.boundTo.(*dualValue)
		if !ok {
			err = NewError(typeError, "Expected a *dualValue in %v (%x). Got %T instead", n, n.Hashcode(), n.boundTo)
			return
		}

		var cached *dualValue
		if cached = s.cache[i]; cached == nil {
			if cached, err = dv.clone0(); err != nil {
				return
			}
			s.cache[i] = cached
		}

		var dt Dtype
		if dt, err = dtypeOf(dv.Type()); err != nil {
			return
		}

		grad := dv.d
		weights := dv.Value

		cv := cached.Value

		switch cw := cv.(type) {
		case Tensor:
			switch dt {
			case Float32:
				var w, g, c, g2, regularized *tf32.Tensor

				l2reg := float32(s.l2reg)
				clip := float32(s.clip)
				eps := float32(s.eps)
				eta := float32(s.eta)

				g = grad.(Tensor).Tensor.(*tf32.Tensor)
				if g2, err = tf32.PointwiseSquare(g); err != nil { // safe version
					return
				}

				c = cw.Tensor.(*tf32.Tensor)
				tf32.Add(c, g2, types.UseUnsafe())

				if s.useClip {
					if _, err = tf32.Clamp(g, -clip, clip, types.UseUnsafe()); err != nil {
						return
					}
				}

				// update
				var upd *tf32.Tensor

				if upd, err = tf32.Add(c, eps); err != nil {
					return
				}

				if _, err = tf32.InvSqrt(upd, types.UseUnsafe()); err != nil {
					return
				}
				if _, err = tf32.PointwiseMul(g, -eta, types.UseUnsafe()); err != nil {
					return
				}

				if _, err = tf32.PointwiseMul(upd, g, types.UseUnsafe()); err != nil {
					return
				}

				// regularize
				w = weights.(Tensor).Tensor.(*tf32.Tensor)

				if s.useL2Reg {
					if regularized, err = tf32.PointwiseMul(w, l2reg); err != nil {
						return
					}

					if _, err = tf32.Sub(upd, regularized, types.UseUnsafe()); err != nil {
						return
					}
				}

				if _, err = tf32.Add(w, upd, types.UseUnsafe()); err != nil {
					return
				}

				// zero all
				g.Zero()
			case Float64:
				var w, g, c, g2, regularized *tf64.Tensor

				l2reg := s.l2reg
				clip := s.clip
				eps := s.eps
				eta := s.eta

				g = grad.(Tensor).Tensor.(*tf64.Tensor)
				if g2, err = tf64.PointwiseSquare(g); err != nil { // safe version
					return
				}

				c = cw.Tensor.(*tf64.Tensor)
				tf64.Add(c, g2, types.UseUnsafe())

				if s.useClip {
					if _, err = tf64.Clamp(g, -clip, clip, types.UseUnsafe()); err != nil {
						return
					}
				}

				// update
				var upd *tf64.Tensor

				if upd, err = tf64.Add(c, eps); err != nil {
					return
				}

				if _, err = tf64.InvSqrt(upd, types.UseUnsafe()); err != nil {
					return
				}
				if _, err = tf64.PointwiseMul(g, -eta, types.UseUnsafe()); err != nil {
					return
				}

				if _, err = tf64.PointwiseMul(upd, g, types.UseUnsafe()); err != nil {
					return
				}

				// regularize
				w = weights.(Tensor).Tensor.(*tf64.Tensor)

				if s.useL2Reg {
					if regularized, err = tf64.PointwiseMul(w, l2reg); err != nil {
						return
					}

					if _, err = tf64.Sub(upd, regularized, types.UseUnsafe()); err != nil {
						return
					}
				}

				if _, err = tf64.Add(w, upd, types.UseUnsafe()); err != nil {
					return
				}

				// zero all
				g.Zero()

			default:

			}
		case Scalar:
			switch dt {
			case Float32:
				var w, g, c float32

				l2reg := float32(s.l2reg)
				clip := float32(s.clip)
				eps := float32(s.eps)
				eta := float32(s.eta)

				c = cw.v.(float32)
				g = grad.(Scalar).v.(float32)

				c += g * g

				if s.useClip {
					if g > clip {
						g = clip
					} else if g < -clip {
						g = -clip
					}
				}

				w = weights.(Scalar).v.(float32)

				upd := -eta * g / math32.Sqrt(c+eps)

				if s.useL2Reg {
					upd -= w * l2reg
				}

				w += upd

				// because scalar values are copies, and not pointers, we have to actually re-update the dualValu in model[i]
				var ws Value
				if ws, err = anyToValue(w); err != nil {
					return
				}

				dv.Value = ws
				dv.d = Scalar{t: Float32, v: float32(0)} // zero it

			case Float64:
				var w, g, c float64

				l2reg := s.l2reg
				clip := s.clip
				eps := s.eps
				eta := s.eta

				c = cw.v.(float64)
				g = grad.(Scalar).v.(float64)

				c += g * g

				if s.useClip {
					if g > clip {
						g = clip
					} else if g < -clip {
						g = -clip
					}
				}

				w = weights.(Scalar).v.(float64)
				upd := -eta * g / math.Sqrt(c+eps)
				if s.useL2Reg {
					upd -= w * l2reg
				}

				w += upd

				// because scalar values are copies, and not pointers, we have to actually re-update the dualValu in model[i]
				var ws Value
				if ws, err = anyToValue(w); err != nil {
					return
				}

				dv.Value = ws
				dv.d = Scalar{t: Float64, v: float64(0)}
			default:
				err = nyi("Adagrad step", dt)
				return
			}
		default:
			err = nyi("Adagrad step", cv)
			return
		}

	}
	return
}
