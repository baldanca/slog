package internal

import (
	"io"
)

type (
	stackModel []string
)

var (
	// stack of logs
	Stack stackModel

	// stackCustom define if Custom logs will be put in stack
	stackCustom bool
	// stackDebug define if Debug logs will be put in stack
	stackDebug bool
	// stackErr define if Err logs will be put in stack
	stackErr bool
	// stackFatal define if Fatal logs will be put in stack
	stackFatal bool
	// stackInfo define if Info logs will be put in stack
	stackInfo bool
	// stackWarn define if Warn logs will be put in stack
	stackWarn bool
	// stackPan define if Panic logs will be put in stack
	stackPan bool
)

func init() {
	verifyStack()
}

// enableAllStack function
func enableAllStack() {
	stackCustom = true
	stackDebug = true
	stackErr = true
	stackFatal = true
	stackInfo = true
	stackWarn = true
	stackPan = true
}

// enableCustomStack function
func enableCustomStack() {
	stackCustom = true
}

// enableDebugStack function
func enableDebugStack() {
	stackDebug = true
}

// enableErrStack function
func enableErrStack() {
	stackErr = true
}

// enableFatalStack function
func enableFatalStack() {
	stackFatal = true
}

// enableInfoStack function
func enableInfoStack() {
	stackInfo = true
}

// enableWarnStack function
func enableWarnStack() {
	stackWarn = true
}

// enablePanStack function
func enablePanStack() {
	stackPan = true
}

// verifyStack function
func verifyStack() {
	if stackDebug || stackErr || stackFatal || stackInfo || stackWarn || stackPan {
		Stack = []string{}
	}
	if stackDebug {
		debugOut = io.MultiWriter(defaultOut, &Stack)
	}
	if stackErr {
		errOut = io.MultiWriter(defaultOut, &Stack)
	}
	if stackFatal {
		fatalOut = io.MultiWriter(defaultOut, &Stack)
	}
	if stackInfo {
		infoOut = io.MultiWriter(defaultOut, &Stack)
	}
	if stackWarn {
		warnOut = io.MultiWriter(defaultOut, &Stack)
	}
	if stackPan {
		panOut = io.MultiWriter(defaultOut, &Stack)
	}
}

// Write a log to stack
func (s *stackModel) Write(p []byte) (int, error) {
	*s = append(*s, string(p))
	return len(p), nil
}
