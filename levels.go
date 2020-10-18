package slog

import (
	"fmt"
	"strings"
)

type (
	// color type
	color string
	// level type
	level string
)

const (
	// Colors
	blue   color = "\033[0;94m%s\033[0m"
	green  color = "\033[0;92m%s\033[0m"
	orange color = "\033[0;38;5;214m%s\033[0m"
	red    color = "\033[0;91m%s\033[0m"
	white  color = "\033[0;0m%s\033[0m"
	yellow color = "\033[0;93m%s\033[0m"
	// Log levels
	customLevel level = "custom"
	debugLevel  level = "debug"
	errLevel    level = "error"
	fatalLevel  level = "fatal"
	infoLevel   level = "info"
	panicLevel  level = "panic"
	warnLevel   level = "warn"
)

var (
	// Log colorize map
	colorizeMap = map[level]color{
		customLevel: white,
		debugLevel:  white,
		errLevel:    red,
		fatalLevel:  yellow,
		infoLevel:   blue,
		panicLevel:  orange,
		warnLevel:   green,
	}
)

// Prefix function
func (l level) Prefix() string {
	prefix := strings.ToUpper(l.String())
	for len(prefix) < 5 {
		prefix += " "
	}
	return fmt.Sprintf("[%s] ", prefix)
}

// Colorize function
func (l level) Colorize() string {
	return fmt.Sprintf(colorizeMap[l].string(), l.Prefix())
}

// String function
func (l level) String() string {
	return string(l)
}

// color to string
func (c color) string() string {
	return string(c)
}
