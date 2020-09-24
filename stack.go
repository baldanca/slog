package slog

import (
	"io"
	"os"
)

type (
	stackModel []string
)

var (
	// stack of logs
	Stack stackModel

	// stackDebug define if Debug logs will be put in stack
	stackDebug bool
	// stackErr define if Err logs will be put in stack
	stackErr bool
	// stackInfo define if Info logs will be put in stack
	stackInfo bool
	// stackWarn define if Warn logs will be put in stack
	stackWarn bool
	// stackPan define if Panic logs will be put in stack
	stackPan bool
)

func init() {
	// Verify environment variables
	if os.Getenv("SLOG_STACK_ALL") == "true" {
		stackDebug = true
		stackErr = true
		stackInfo = true
		stackWarn = true
		stackPan = true
	}
	if os.Getenv("SLOG_STACK_DEBUG") == "true" {
		stackDebug = true
	}
	if os.Getenv("SLOG_STACK_ERROR") == "true" {
		stackErr = true
	}
	if os.Getenv("SLOG_STACK_INFO") == "true" {
		stackInfo = true
	}
	if os.Getenv("SLOG_STACK_WARN") == "true" {
		stackWarn = true
	}
	if os.Getenv("SLOG_STACK_PANIC") == "true" {
		stackPan = true
	}
	enableStack()
}

// enableStack function
func enableStack() {
	if stackDebug || stackErr || stackInfo || stackWarn || stackPan {
		Stack = []string{}
	}
	if stackDebug {
		debugOut = io.MultiWriter(defaultOut, &Stack)
	}
	if stackErr {
		errOut = io.MultiWriter(defaultOut, &Stack)
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
	updateLoggers()
}

// Write a log to stack
func (s *stackModel) Write(p []byte) (int, error) {
	*s = append(*s, string(p))
	return len(p), nil
}
