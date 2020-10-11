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
	// Setting standard
	SetStdLog(New(os.Stdout, log.Ldate|log.Ltime).NoLog().Humanize())
}

// SetStdLog function
func SetStdLog(l Service) {
	StdLog = l
}

// AddHandler to standard logger
func AddHandler(h Handler) {
	StdLog.AddHandler(h)
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

/*
// Panic logger
func (l *Logger) Panic(v ...interface{}) {
	v = l.loadHandlers(v...)
	l.panic.Output(l.calldepth, fmt.Sprintln(v...))
	panic(fmt.Sprint(v...))
}

// Panicf logger with format
func (l *Logger) Panicf(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}
	v = l.loadHandlers(v...)
	l.panic.Output(l.calldepth, fmt.Sprintf(format, v...))
	panic(fmt.Sprint(v...))
}

// Warn logger
func (l *Logger) Warn(v ...interface{}) {
	v = l.loadHandlers(v...)
	l.warn.Output(l.calldepth, fmt.Sprintln(v...))
}

// Warnf logger with format
func (l *Logger) Warnf(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}
	v = l.loadHandlers(v...)
	l.warn.Output(l.calldepth, fmt.Sprintf(format, v...))
}

*/
