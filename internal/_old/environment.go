package internal

import (
	"os"
)

var (
	// Environment mapping
	mapEnv = map[string]func(){
		// Debug environment flag
		"SLOG_DEBUG": enableDebug,

		// Humanize environment flag
		"SLOG_HUMANIZE": enableHumanize,

		// Colorize environment flag
		"SLOG_COLORIZE": enableColorize,

		// Stack environment flag
		"SLOG_STACK_ALL":    enableAllStack,
		"SLOG_STACK_CUSTOM": enableCustomStack,
		"SLOG_STACK_DEBUG":  enableDebugStack,
		"SLOG_STACK_ERROR":  enableErrStack,
		"SLOG_STACK_INFO":   enableInfoStack,
		"SLOG_STACK_WARN":   enableWarnStack,
		"SLOG_STACK_PANIC":  enablePanStack,
	}
)

func init() {
	// Verifying environment variables
	for env, fun := range mapEnv {
		if os.Getenv(env) == "true" {
			fun()
		}
	}
}
