package handler

import (
	"errors"
	"testing"

	"github.com/luigiBaldanza/slog/test"
)

func TestHandler(t *testing.T) {
	var (
		stringTest  = "TEST"
		pointerTest = &stringTest
		intTest     = 369
		errorTest   = errors.New("TEST")
		jsonTest    = []byte(`{"teste01":"TEST","teste02":"TEST","teste03":369,"teste04":"TEST"}`)
		funcTest    = func(teste string) bool { return true }
		mapTest     = map[string]interface{}{
			"teste01": stringTest,
			"teste02": pointerTest,
			"teste03": intTest,
			"teste04": errorTest,
			"teste05": jsonTest,
			"teste06": funcTest,
		}
		arrayTest = [7]interface{}{
			stringTest,
			pointerTest,
			intTest,
			errorTest,
			jsonTest,
			funcTest,
			mapTest,
		}
		sliceTest = []interface{}{
			stringTest,
			pointerTest,
			intTest,
			errorTest,
			jsonTest,
			funcTest,
			mapTest,
			arrayTest,
		}
	)
	handlers := NewHandlers()
	handlers.Add(func(values ...interface{}) []interface{} {
		return values
	})
	test.Assert(t, sliceTest, handlers.Run(sliceTest...))
}
