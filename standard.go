package slog

import (
	"log"
	"os"
)

var (
	// StdLog is a standard log configuration
	StdLog Service
)

func init() {
	// Setting standard logger
	SetStdLog(New(os.Stdout, log.Ldate|log.Ltime).NoLog().Humanize())
}

// SetStdLog function
func SetStdLog(l Service) {
	StdLog = l
}

// Custom standard logger
func Custom(calldepth int, prefix string, v ...interface{}) {
	StdLog.Custom(calldepth, prefix, v...)
}

// Customf standard logger with format
func Customf(calldepth int, prefix, format string, v ...interface{}) {
	StdLog.Customf(calldepth, prefix, format, v...)
}

// Debug standard logger
func Debug(v ...interface{}) {
	StdLog.Debug(v...)
}

// Debugf standard logger with format
func Debugf(format string, v ...interface{}) {
	StdLog.Debugf(format, v...)
}

// Error standard logger
func Error(v ...interface{}) {
	StdLog.Error(v...)
}

// Errorf standard logger with format
func Errorf(format string, v ...interface{}) {
	StdLog.Errorf(format, v...)
}

// Fatal standard logger
func Fatal(v ...interface{}) {
	StdLog.Fatal(v...)
}

// Fatalf standard logger with format
func Fatalf(format string, v ...interface{}) {
	StdLog.Fatalf(format, v...)
}

// Info standard logger
func Info(v ...interface{}) {
	StdLog.Info(v...)
}

// Infof standard logger with format
func Infof(format string, v ...interface{}) {
	StdLog.Infof(format, v...)
}

// Panic standard logger
func Panic(v ...interface{}) {
	StdLog.Panic(v...)
}

// Panicf standard logger with format
func Panicf(format string, v ...interface{}) {
	StdLog.Panicf(format, v...)
}

// Warn standard logger
func Warn(v ...interface{}) {
	StdLog.Warn(v...)
}

// Warnf standard logger with format
func Warnf(format string, v ...interface{}) {
	StdLog.Warnf(format, v...)
}

// EnableDebug on standard logger
func EnableDebug() Service {
	StdLog.EnableDebug()
	return StdLog
}

// Colorize on standard logger
func Colorize() Service {
	StdLog.Colorize()
	return StdLog
}

// AddHandler on standard logger
func AddHandler(h Handler) Service {
	StdLog.AddHandler(h)
	return StdLog
}

// Humanize on standard logger
func Humanize() Service {
	StdLog.Humanize()
	return StdLog
}

// NoLog on standard logger
func NoLog() Service {
	StdLog.NoLog()
	return StdLog
}

// NewStack on standard logger
func NewStack() Service {
	StdLog.NewStack()
	return StdLog
}
