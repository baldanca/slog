package slog

import (
	"strings"

	"github.com/luigiBaldanza/slog/level"
)

type (
	// Logger interface type
	Logger interface {
		Log(calldepth int, logLevel level.Type, values ...interface{})
		Debug(values ...interface{})
		Info(values ...interface{})
		Warn(values ...interface{})
		Error(values ...interface{})
		Fatal(values ...interface{})
		Panic(values ...interface{})
		/*

			Logf(calldepth int, prefix, format string, values ...interface{})
			Debugf(format string, values ...interface{})
			Infof(format string, values ...interface{})
			Warnf(format string, values ...interface{})
			Errorf(format string, values ...interface{})
			Fatalf(format string, values ...interface{})
			Panicf(format string, values ...interface{})

			SetDefaultOut(out io.Writer)
			SetCustomOut(out io.Writer)
			SetDebugOut(out io.Writer)
			SetErrorOut(out io.Writer)
			SetFatalOut(out io.Writer)
			SetInfoOut(out io.Writer)
			SetPanicOut(out io.Writer)
			SetWarnOut(out io.Writer)

			EnableDebug() */
		// 	Next implementations:
		// 		Colorize()

		// 		AddHandler(h handler.Func)
		// 		Humanize()
		// 		NoLog()

		// 		Tracef(format string, args ...interface{})
		// 		Trace(args ...interface{})
		// 		Traceln(args ...interface{})
	}
)

// verifyFormat input logger
func verifyFormat(format *string) {
	if !strings.HasSuffix(*format, "\n") {
		*format = *format + "\n"
	}
}

/*

// verifyFlags input logger
func (s *Slog) verifyFlags() {
	if s.noLog {
		l.AddHandler(noLogAll)
	}
	if s.humanize {
		l.AddHandler(humanizeAll)
	}
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
*/
