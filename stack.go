package slog

import "io"

type (
	// Stack model
	Stack struct {
		data [][]byte
		out  io.Writer
	}
)

/*
// EnableStack function
func (l *Logger) EnableStack() *Logger {
	l.custom.SetOutput(io.MultiWriter(l.custom.Writer(), l.stack))
	l.debug.SetOutput(io.MultiWriter(l.debug.Writer(), l.stack))
	l.err.SetOutput(io.MultiWriter(l.err.Writer(), l.stack))
	l.fatal.SetOutput(io.MultiWriter(l.fatal.Writer(), l.stack))
	l.info.SetOutput(io.MultiWriter(l.info.Writer(), l.stack))
	l.panic.SetOutput(io.MultiWriter(l.panic.Writer(), l.stack))
	l.warn.SetOutput(io.MultiWriter(l.warn.Writer(), l.stack))
	return l
}

// GetAll stack function
func (s stack) GetAll() stack {
	return s
}

// GetAllString stack function
func (s stack) GetAllString() []string {
	stackString := []string{}
	for _, value := range s {
		stackString = append(stackString, string(value))
	}
	return stackString
}

// PrintAll stack function
func (s stack) PrintAll() {
	for _, value := range s {
		fmt.Printf(string(value))
	}
}

*/

// Write a log to stack
func (s Stack) Write(p []byte) (int, error) {
	s.data = append(s.data, p)
	return len(p), nil
}
