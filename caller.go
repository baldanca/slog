package slog

import (
	"fmt"
	"runtime"
	"strings"
)

func caller(calldepth int) string {
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	path := strings.Split(file, "/")
	file = path[len(path)-1]
	return fmt.Sprintf("%s:%d", file, line)
}
