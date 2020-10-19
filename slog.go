package slog

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	// ErrInvalidPrefix error
	ErrInvalidPrefix = errors.New("invalid prefix")
	// ErrPrefixTooLong error
	ErrPrefixTooLong = errors.New("prefix too long")
)

type (
	// LoggerService contract
	LoggerService interface {
		// Loggers
		Custom(calldepth int, prefix string, v ...interface{})
		Customf(calldepth int, prefix, format string, v ...interface{})
		Debug(v ...interface{})
		Debugf(format string, v ...interface{})
		Error(v ...interface{})
		Errorf(format string, v ...interface{})
		Fatal(v ...interface{})
		Fatalf(format string, v ...interface{})
		Info(v ...interface{})
		Infof(format string, v ...interface{})
		Panic(v ...interface{})
		Panicf(format string, v ...interface{})
		Warn(v ...interface{})
		Warnf(format string, v ...interface{})
		// Outputs
		SetDefaultOut(out io.Writer) *Logger
		SetCustomOut(out io.Writer) *Logger
		SetDebugOut(out io.Writer) *Logger
		SetErrorOut(out io.Writer) *Logger
		SetFatalOut(out io.Writer) *Logger
		SetInfoOut(out io.Writer) *Logger
		SetPanicOut(out io.Writer) *Logger
		SetWarnOut(out io.Writer) *Logger
		// Features
		EnableDebug() *Logger
		Colorize() *Logger
		// Handler
		AddHandler(h Handler) *Logger
		Humanize() *Logger
		NoLog() *Logger
	}
	// Logger model
	Logger struct {
		calldepth int
		handlers  *Handlers
		noLog     bool
		humanize  bool
		stack     *Stack
		// file      *File
		custom *log.Logger
		debug  *log.Logger
		err    *log.Logger
		fatal  *log.Logger
		info   *log.Logger
		panic  *log.Logger
		warn   *log.Logger
	}
)

// New logger service
func New(out io.Writer, flag int) LoggerService {
	return &Logger{
		calldepth: 2,
		handlers:  NewHandlers(),
		noLog:     false,
		humanize:  false,
		custom:    log.New(out, customLevel.Prefix(), flag),
		debug:     log.New(ioutil.Discard, debugLevel.Prefix(), flag),
		err:       log.New(out, errLevel.Prefix(), flag),
		fatal:     log.New(out, fatalLevel.Prefix(), flag),
		info:      log.New(out, infoLevel.Prefix(), flag),
		panic:     log.New(out, panicLevel.Prefix(), flag),
		warn:      log.New(out, warnLevel.Prefix(), flag),
	}
}

// verifyPrefix input logger
func (l *Logger) verifyPrefix(prefix string) {
	if prefix == "" {
		panic(ErrInvalidPrefix)
	}
	if len(prefix) > 5 {
		panic(ErrPrefixTooLong)
	}
}

// verifyFormat input logger
func (l *Logger) verifyFormat(format *string) {
	if !strings.HasSuffix(*format, "\n") {
		*format = *format + "\n"
	}
}

// verifyFlags input logger
func (l *Logger) verifyFlags() {
	if l.noLog {
		l.AddHandler(noLogAll)
	}
	if l.humanize {
		l.AddHandler(humanizeAll)
	}
}

// Custom logger
func (l *Logger) Custom(calldepth int, prefix string, v ...interface{}) {
	l.verifyPrefix(prefix)
	l.custom.SetPrefix(level(prefix).Prefix())
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.custom.Output(calldepth, fmt.Sprintln(v...))
}

// Customf logger with format
func (l *Logger) Customf(calldepth int, prefix, format string, v ...interface{}) {
	l.verifyPrefix(prefix)
	l.custom.SetPrefix(level(prefix).Prefix())
	l.verifyFormat(&format)
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.custom.Output(calldepth, fmt.Sprintf(format, v...))
}

// Debug logger
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Output(l.calldepth, fmt.Sprintln(v...))
}

// Debugf logger with format
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.verifyFormat(&format)
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.debug.Output(l.calldepth, fmt.Sprintf(format, v...))
}

// Error logger
func (l *Logger) Error(v ...interface{}) {
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.err.Output(l.calldepth, fmt.Sprintln(v...))
}

// Errorf logger with format
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.verifyFormat(&format)
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.err.Output(l.calldepth, fmt.Sprintf(format, v...))
}

// Fatal logger
func (l *Logger) Fatal(v ...interface{}) {
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.fatal.Output(l.calldepth, fmt.Sprintln(v...))
	os.Exit(1)
}

// Fatalf logger with format
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.verifyFormat(&format)
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.fatal.Output(l.calldepth, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Info logger
func (l *Logger) Info(v ...interface{}) {
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.info.Output(l.calldepth, fmt.Sprintln(v...))
}

// Infof logger with format
func (l *Logger) Infof(format string, v ...interface{}) {
	l.verifyFormat(&format)
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.info.Output(l.calldepth, fmt.Sprintf(format, v...))
}

// Panic logger
func (l *Logger) Panic(v ...interface{}) {
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.panic.Output(l.calldepth, fmt.Sprintln(v...))
	panic(fmt.Sprint(v...))
}

// Panicf logger with format
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.verifyFormat(&format)
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.panic.Output(l.calldepth, fmt.Sprintf(format, v...))
	panic(fmt.Sprint(v...))
}

// Warn logger
func (l *Logger) Warn(v ...interface{}) {
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.warn.Output(l.calldepth, fmt.Sprintln(v...))
}

// Warnf logger with format
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.verifyFormat(&format)
	l.verifyFlags()
	v = l.handlers.run(v...)
	l.warn.Output(l.calldepth, fmt.Sprintf(format, v...))
}

// SetDefaultOut logger
func (l *Logger) SetDefaultOut(out io.Writer) *Logger {
	l.custom.SetOutput(out)
	l.debug.SetOutput(out)
	l.err.SetOutput(out)
	l.fatal.SetOutput(out)
	l.info.SetOutput(out)
	l.panic.SetOutput(out)
	l.warn.SetOutput(out)
	return l
}

// SetCustomOut logger
func (l *Logger) SetCustomOut(out io.Writer) *Logger {
	l.custom.SetOutput(out)
	return l
}

// SetDebugOut logger
func (l *Logger) SetDebugOut(out io.Writer) *Logger {
	l.debug.SetOutput(out)
	return l
}

// SetErrorOut logger
func (l *Logger) SetErrorOut(out io.Writer) *Logger {
	l.err.SetOutput(out)
	return l
}

// SetFatalOut logger
func (l *Logger) SetFatalOut(out io.Writer) *Logger {
	l.fatal.SetOutput(out)
	return l
}

// SetInfoOut logger
func (l *Logger) SetInfoOut(out io.Writer) *Logger {
	l.info.SetOutput(out)
	return l
}

// SetPanicOut logger
func (l *Logger) SetPanicOut(out io.Writer) *Logger {
	l.panic.SetOutput(out)
	return l
}

// SetWarnOut logger
func (l *Logger) SetWarnOut(out io.Writer) *Logger {
	l.warn.SetOutput(out)
	return l
}

// EnableDebug function
func (l *Logger) EnableDebug() *Logger {
	l.debug.SetOutput(os.Stdout)
	return l
}

// Colorize function
func (l *Logger) Colorize() *Logger {
	l.custom.SetPrefix(customLevel.Colorize())
	l.debug.SetPrefix(debugLevel.Colorize())
	l.err.SetPrefix(errLevel.Colorize())
	l.fatal.SetPrefix(fatalLevel.Colorize())
	l.info.SetPrefix(infoLevel.Colorize())
	l.panic.SetPrefix(panicLevel.Colorize())
	l.warn.SetPrefix(warnLevel.Colorize())
	return l
}

// AddHandler logger
func (l *Logger) AddHandler(h Handler) *Logger {
	l.handlers.Add(h)
	return l
}

// Humanize function
func (l *Logger) Humanize() *Logger {
	l.humanize = true
	return l
}

// NoLog function
func (l *Logger) NoLog() *Logger {
	l.noLog = true
	return l
}
