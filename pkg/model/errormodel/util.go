package errormodel

import "runtime"

type TraceInfo struct {
	Name string
	File string
	Line int
}

func callers() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(4, pcs[:])
	return pcs[0:n]
}

func getStackTrace() []TraceInfo {
	stackTrace := make([]TraceInfo, 0, 32)
	callStack := callers()
	for k := range callStack {
		v := callStack[k] - 1
		f := runtime.FuncForPC(v)
		file, line := f.FileLine(v)

		stackTrace = append(stackTrace,
			TraceInfo{
				Name: f.Name(),
				File: file,
				Line: line,
			})
	}
	return stackTrace
}
