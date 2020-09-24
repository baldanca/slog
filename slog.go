package slog

import (
	"io"
	"log"
	"os"
)

const (
	flags = log.Ldate | log.Ltime
)

var (
	// debugFlag define if debug is enabled
	debugFlag bool

	// debug logger
	debug *log.Logger
	// error logger
	err *log.Logger
	// info logger
	info *log.Logger
	// warn logger
	warn *log.Logger
	// panic logger
	pan *log.Logger

	// default output of logs
	defaultOut io.Writer

	// output of debug logger
	debugOut io.Writer
	// output of err logger
	errOut io.Writer
	// output of info logger
	infoOut io.Writer
	// output of warn logger
	warnOut io.Writer
	// output of panic logger
	panOut io.Writer
)

func init() {
	// Define default output
	defaultOut = os.Stdout

	// Verify environment variables
	if os.Getenv("SLOG_DEBUG") == "true" {
		enableDebug()
	}

	setOutputs(defaultOut)
	updateLoggers()
}

func setOutputs(out io.Writer) {
	debugOut = out
	errOut = out
	infoOut = out
	warnOut = out
	panOut = out
}

func updateLoggers() {
	debug = log.New(debugOut, "[DEBUG] ", flags)
	err = log.New(errOut, "[ERROR] ", flags)
	info = log.New(infoOut, "[INFO ] ", flags)
	warn = log.New(warnOut, "[WARN ] ", flags)
	pan = log.New(panOut, "[PANIC] ", flags)
}

// enableDebug function
func enableDebug() {
	debugFlag = true
}

// Custom function
func Custom(prefix string, calldepth int, i ...interface{}) {
	log.New(defaultOut, prefix+" ", flags).Println(humanizeAll(append([]interface{}{caller(calldepth)}, i...)...)...)
}

// Debug function
func Debug(i ...interface{}) {
	if debugFlag {
		debug.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
	}
}

// Error function
func Error(i ...interface{}) {
	err.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
}

// Info function
func Info(i ...interface{}) {
	info.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
}

// Panic function
func Panic(i ...interface{}) {
	pan.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
}

// Warn function
func Warn(i ...interface{}) {
	warn.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
}
