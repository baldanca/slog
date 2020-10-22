package slog

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/luigiBaldanza/slog/level"
)

type (
	// Slog struct type
	Slog struct {
		log   *log.Logger
		debug *log.Logger
		info  *log.Logger
		warn  *log.Logger
		err   *log.Logger
		fatal *log.Logger
		panic *log.Logger
	}
)

// New logger service
func New(out io.Writer, flag int) Logger {
	return &Slog{
		log:   log.New(out, level.Log.Prefix(), flag),
		debug: log.New(ioutil.Discard, level.Debug.Prefix(), flag),
		err:   log.New(out, level.Err.Prefix(), flag),
		fatal: log.New(out, level.Fatal.Prefix(), flag),
		info:  log.New(out, level.Info.Prefix(), flag),
		panic: log.New(out, level.Panic.Prefix(), flag),
		warn:  log.New(out, level.Warn.Prefix(), flag),
	}
}

// Log logger
func (s *Slog) Log(calldepth int, logLevel level.Type, values ...interface{}) {
	logLevel.ValidatePrefix()
	s.log.SetPrefix(logLevel.Prefix())
	s.log.Output(calldepth, fmt.Sprintln(values...))
}

// Debug logger
func (s *Slog) Debug(values ...interface{}) {
	s.Log(2, level.Debug, values...)
}

// Info logger
func (s *Slog) Info(values ...interface{}) {
	s.Log(2, level.Info, values...)
}

// Warn logger
func (s *Slog) Warn(values ...interface{}) {
	s.Log(2, level.Warn, values...)
}

// Error logger
func (s *Slog) Error(values ...interface{}) {
	s.Log(2, level.Err, values...)
}

// Fatal logger
func (s *Slog) Fatal(values ...interface{}) {
	s.Log(2, level.Fatal, values...)
	os.Exit(1)
}

// Panic logger
func (s *Slog) Panic(values ...interface{}) {
	s.Log(2, level.Panic, values...)
	panic(values)
}
