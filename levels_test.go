package slog

import (
	"fmt"
	"testing"
)

func TestLevelPrefix(t *testing.T) {
	// Custom
	err := test(t, "[CUSTOM] ", customLevel.Prefix())
	if err != nil {
		t.Error(err)
	}

	// Debug
	err = test(t, "[DEBUG] ", debugLevel.Prefix())
	if err != nil {
		t.Error(err)
	}

	// Error
	err = test(t, "[ERROR] ", errLevel.Prefix())
	if err != nil {
		t.Error(err)
	}

	// Fatal
	err = test(t, "[FATAL] ", fatalLevel.Prefix())
	if err != nil {
		t.Error(err)
	}

	// Info
	err = test(t, "[INFO ] ", infoLevel.Prefix())
	if err != nil {
		t.Error(err)
	}

	// Panic
	err = test(t, "[PANIC] ", panicLevel.Prefix())
	if err != nil {
		t.Error(err)
	}

	// Warn
	err = test(t, "[WARN ] ", warnLevel.Prefix())
	if err != nil {
		t.Error(err)
	}
}

func TestLevelColor(t *testing.T) {
	// White / Custom
	err := test(t, fmt.Sprintf(white.string(), customLevel.Prefix()), customLevel.Colorize())
	if err != nil {
		t.Error(err)
	}

	// White / Debug
	err = test(t, fmt.Sprintf(white.string(), debugLevel.Prefix()), debugLevel.Colorize())
	if err != nil {
		t.Error(err)
	}

	// Red / Error
	err = test(t, fmt.Sprintf(red.string(), errLevel.Prefix()), errLevel.Colorize())
	if err != nil {
		t.Error(err)
	}

	// Yellow / Fatal
	err = test(t, fmt.Sprintf(yellow.string(), fatalLevel.Prefix()), fatalLevel.Colorize())
	if err != nil {
		t.Error(err)
	}

	// Blue / Info
	err = test(t, fmt.Sprintf(blue.string(), infoLevel.Prefix()), infoLevel.Colorize())
	if err != nil {
		t.Error(err)
	}

	// Orange / Panic
	err = test(t, fmt.Sprintf(orange.string(), panicLevel.Prefix()), panicLevel.Colorize())
	if err != nil {
		t.Error(err)
	}

	// Green / Warn
	err = test(t, fmt.Sprintf(green.string(), warnLevel.Prefix()), warnLevel.Colorize())
	if err != nil {
		t.Error(err)
	}
}
