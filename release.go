// +build !debug

package gorgonia

// DEBUG indicates if this build is in debug mode. It is not.
const DEBUG = false

const (
	compileDev        = false
	shapeInferenceDev = false
	typeSystemDev     = false
	symdiffDev        = false
	autodiffDev       = false
	machineDev        = false
	stabilizationDev  = false
	cudaDev           = false
)

func tabcount() int { return 0 }

func enterLoggingContext()                                    {}
func leaveLoggingContext()                                    {}
func logf(format string, others ...interface{})               {}
func compileLogf(format string, attrs ...interface{})         {}
func shapeLogf(format string, attrs ...interface{})           {}
func typeSysLogf(format string, attrs ...interface{})         {}
func symdiffLogf(format string, attrs ...interface{})         {}
func autodiffLogf(format string, attrs ...interface{})        {}
func machineLogf(format string, attrs ...interface{})         {}
func stabLogf(format string, attrs ...interface{})            {}
func solverLogf(format string, attrs ...interface{})          {}
func cudaLogf(format string, attrs ...interface{})            {}
func recoverFrom(format string, attrs ...interface{})         {}
func logCompileState(name string, g *ExprGraph, df *dataflow) {}

func incrCC() {}
func incrEC() {}
func incrNN() {}
