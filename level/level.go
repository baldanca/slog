package level

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrInvalidPrefix error
	ErrInvalidPrefix = errors.New("invalid prefix")
	// ErrPrefixTooLong error
	ErrPrefixTooLong = errors.New("prefix too long")
)

type (
	// colorType string type
	colorType string
	// Type string type
	Type string
)

const (
	blue   colorType = "\033[0;94m%s\033[0m"
	green  colorType = "\033[0;92m%s\033[0m"
	orange colorType = "\033[0;38;5;214m%s\033[0m"
	red    colorType = "\033[0;91m%s\033[0m"
	white  colorType = "\033[0;0m%s\033[0m"
	yellow colorType = "\033[0;93m%s\033[0m"

	// Log level
	Log Type = "log"
	// Debug level
	Debug Type = "debug"
	// Err level
	Err Type = "error"
	// Fatal level
	Fatal Type = "fatal"
	// Info level
	Info Type = "info"
	// Panic level
	Panic Type = "panic"
	// Warn level
	Warn Type = "warn"
)

var (
	colorizeMap = map[Type]colorType{
		Log:   white,
		Debug: white,
		Err:   red,
		Fatal: yellow,
		Info:  blue,
		Panic: orange,
		Warn:  green,
	}
)

// NewPrefix factory
func NewPrefix(prefix string) Type {
	return Type(prefix)
}

// ValidatePrefix input logger
func (l Type) ValidatePrefix() {
	if l == "" {
		panic(ErrInvalidPrefix)
	}
	if len(l) > 5 {
		panic(ErrPrefixTooLong)
	}
}

// Prefix function
func (l Type) Prefix() string {
	prefix := strings.ToUpper(l.String())
	for len(prefix) < 5 {
		prefix += " "
	}
	return fmt.Sprintf("[%s] ", prefix)
}

// Colorize function
func (l Type) Colorize() string {
	return fmt.Sprintf(colorizeMap[l].string(), l.Prefix())
}

// String function
func (l Type) String() string {
	return string(l)
}

// color to string
func (c colorType) string() string {
	return string(c)
}
