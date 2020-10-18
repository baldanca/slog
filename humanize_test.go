package slog

import (
	"testing"
)

func TestHumanizeAll(t *testing.T) {
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
	for i, result := range humanizeAll(expected...) {
		err := test(t, expected[i], result)
		if err != nil {
			t.Error(err)
		}
	}
}
