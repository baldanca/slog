package handler

import (
	"errors"
	"testing"

	"github.com/luigiBaldanza/slog/test"
)

func TestNoLogAll(t *testing.T) {
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
		structTeste = test.DataTest{
			String:  stringTest,
			Pointer: pointerTest,
			Int:     intTest,
			Error:   errorTest,
			JSON:    jsonTest,
			Func:    funcTest,
			Map:     mapTest,
			Array:   arrayTest,
			Slice:   sliceTest,
			Logged:  stringTest,
		}
		structTesteWithNil = test.DataTest{
			String:  "",
			Pointer: nil,
			Int:     0,
			Error:   nil,
			JSON:    nil,
			Func:    nil,
			Map:     nil,
			Array:   [7]interface{}{},
			Slice:   nil,
			Logged:  "",
		}
	)
	expected := []interface{}{
		stringTest,
		pointerTest,
		intTest,
		errorTest,
		jsonTest,
		funcTest,
		mapTest,
		arrayTest,
		sliceTest,
		structTeste,
		structTesteWithNil,
		nil,
	}
	for index, result := range NoLogAll(expected...) {
		test.Assert(t, expected[index], result)
	}
}
