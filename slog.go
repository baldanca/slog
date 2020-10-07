package slog

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type (
	// service contract
	Service interface {
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

		EnableColorize() *Logger
		EnableDebug() *Logger
		EnableHumanize() *Logger
		// EnableNoLog() *Logger
		// EnableStack() *Logger

		// LoadEnvironment() *Logger

		AddHandler(h Handler)

		// SETOUTPUT
	}

	// handler function
	Handler func(v ...interface{}) []interface{}

	// Logger model
	Logger struct {
		calldepth int
		handlers  []Handler
		humanize  bool
		nolog     bool
		custom    *log.Logger
		debug     *log.Logger
		err       *log.Logger
		fatal     *log.Logger
		info      *log.Logger
		panic     *log.Logger
		warn      *log.Logger
	}
)

// New logger service
func New(out io.Writer, flag int) Service {
	return &Logger{
		calldepth: 2,
		handlers:  []Handler{},
		humanize:  false,
		nolog:     false,
		custom:    log.New(out, customLevel.Prefix(), flag),
		debug:     log.New(ioutil.Discard, debugLevel.Prefix(), flag),
		err:       log.New(out, errLevel.Prefix(), flag),
		fatal:     log.New(out, fatalLevel.Prefix(), flag),
		info:      log.New(out, infoLevel.Prefix(), flag),
		panic:     log.New(out, panicLevel.Prefix(), flag),
		warn:      log.New(out, warnLevel.Prefix(), flag),
	}
}

// SetStdLog function
func SetStdLog(l Service) {
	StdLog = l
}

// AddHandler logger
func (l *Logger) AddHandler(h Handler) {
	l.handlers = append(l.handlers, h)
}

// loadHandlers function
func (l *Logger) loadHandlers(v ...interface{}) []interface{} {
	for _, handler := range l.handlers {
		v = handler(v...)
	}
	return v
}

// Custom logger
func (l *Logger) Custom(calldepth int, prefix string, v ...interface{}) {
	if prefix == "" {
		panic("invalid custom prefix")
	}
	if len(prefix) > 5 {
		panic("custom prefix is too long")
	}
	l.custom.SetPrefix(level(prefix).Prefix())
	v = l.loadHandlers(v...)
	l.custom.Output(calldepth, fmt.Sprintln(v...))
}

// Customf logger with format
func (l *Logger) Customf(calldepth int, prefix, format string, v ...interface{}) {
	if prefix == "" {
		panic("invalid custom prefix")
	}
	if len(prefix) > 5 {
		panic("custom prefix is too long")
	}
	l.custom.SetPrefix(level(prefix).Prefix())
	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}
	v = l.loadHandlers(v...)
	l.custom.Output(calldepth, fmt.Sprintf(format, v...))
}

// Debug logger
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Output(l.calldepth, fmt.Sprintln(v...))
}

// Debugf logger with format
func (l *Logger) Debugf(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}
	v = l.loadHandlers(v...)
	l.debug.Output(l.calldepth, fmt.Sprintf(format, v...))
}

// Error logger
func (l *Logger) Error(v ...interface{}) {
	v = l.loadHandlers(v...)
	l.err.Output(l.calldepth, fmt.Sprintln(v...))
}

// Errorf logger with format
func (l *Logger) Errorf(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}
	v = l.loadHandlers(v...)
	l.err.Output(l.calldepth, fmt.Sprintf(format, v...))
}

// Fatal logger
func (l *Logger) Fatal(v ...interface{}) {
	v = l.loadHandlers(v...)
	l.fatal.Output(l.calldepth, fmt.Sprintln(v...))
	os.Exit(1)
}

// Fatalf logger with format
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}
	v = l.loadHandlers(v...)
	l.fatal.Output(l.calldepth, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Info logger
func (l *Logger) Info(v ...interface{}) {
	v = l.loadHandlers(v...)
	l.info.Output(l.calldepth, fmt.Sprintln(v...))
}

// Infof logger with format
func (l *Logger) Infof(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}
	v = l.loadHandlers(v...)
	l.info.Output(l.calldepth, fmt.Sprintf(format, v...))
}

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

// EnableColorize function
func (l *Logger) EnableColorize() *Logger {
	l.custom.SetPrefix(customLevel.Colorize())
	l.debug.SetPrefix(debugLevel.Colorize())
	l.err.SetPrefix(errLevel.Colorize())
	l.fatal.SetPrefix(fatalLevel.Colorize())
	l.info.SetPrefix(infoLevel.Colorize())
	l.panic.SetPrefix(panicLevel.Colorize())
	l.warn.SetPrefix(warnLevel.Colorize())
	return l
}

// EnableDebug function
func (l *Logger) EnableDebug() *Logger {
	l.debug.SetOutput(os.Stdout)
	return l
}

// EnableHumanize function
func (l *Logger) EnableHumanize() *Logger {
	l.AddHandler(humanizeAll)
	return l
}

/*
// EnableNoLog function
func (l *Logger) EnableNoLog() *Logger {
	return l
}

// EnableStack function
func (l *Logger) EnableStack() *Logger {
	return l
}

// LoadEnvironment function
func (l *Logger) LoadEnvironment() *Logger {
	return l
}
*/
