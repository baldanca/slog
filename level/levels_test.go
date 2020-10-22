package level

import (
	"fmt"
	"testing"

	"github.com/luigiBaldanza/slog/test"
)

func TestLevelPrefix(t *testing.T) {
	// Custom
	test.Assert(t, "[CUSTOM] ", Log.Prefix())

	// Debug
	test.Assert(t, "[DEBUG] ", Debug.Prefix())

	// Error
	test.Assert(t, "[ERROR] ", Err.Prefix())

	// Fatal
	test.Assert(t, "[FATAL] ", Fatal.Prefix())

	// Info
	test.Assert(t, "[INFO ] ", Info.Prefix())

	// Panic
	test.Assert(t, "[PANIC] ", Panic.Prefix())

	// Warn
	test.Assert(t, "[WARN ] ", Warn.Prefix())
}

func TestLevelColor(t *testing.T) {
	// White / Custom
	test.Assert(t, fmt.Sprintf(white.string(), Log.Prefix()), Log.Colorize())

	// White / Debug
	test.Assert(t, fmt.Sprintf(white.string(), Debug.Prefix()), Debug.Colorize())

	// Red / Error
	test.Assert(t, fmt.Sprintf(red.string(), Err.Prefix()), Err.Colorize())

	// Yellow / Fatal
	test.Assert(t, fmt.Sprintf(yellow.string(), Fatal.Prefix()), Fatal.Colorize())

	// Blue / Info
	test.Assert(t, fmt.Sprintf(blue.string(), Info.Prefix()), Info.Colorize())

	// Orange / Panic
	test.Assert(t, fmt.Sprintf(orange.string(), Panic.Prefix()), Panic.Colorize())

	// Green / Warn
	test.Assert(t, fmt.Sprintf(green.string(), Warn.Prefix()), Warn.Colorize())
}
