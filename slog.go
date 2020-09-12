package slog

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	flags = log.Ldate | log.Ltime

	// debug logger
	debug *log.Logger
	// error logger
	err *log.Logger
	// info logger
	info *log.Logger
	// warn logger
	warn *log.Logger
	// panic logger
	panic *log.Logger
)

func init() {
	debug = log.New(ioutil.Discard, "[DEBUG] ", flags)
	err = log.New(os.Stdout, "[ERROR] ", flags)
	info = log.New(os.Stdout, "[INFO ] ", flags)
	warn = log.New(os.Stdout, "[WARN ] ", flags)
	panic = log.New(os.Stdout, "[PANIC] ", flags)
}

// EnableDebug function
func EnableDebug() {
	debug.SetOutput(os.Stdout)
}

// Debug function
func Debug(i ...interface{}) {
	debug.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
}

// Error function
func Error(i ...interface{}) {
	err.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
}

// Info function
func Info(i ...interface{}) {
	info.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
}

// Warn function
func Warn(i ...interface{}) {
	warn.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
}

// Panic function
func Panic(i ...interface{}) {
	panic.Println(humanizeAll(append([]interface{}{caller(2)}, i...)...)...)
}
