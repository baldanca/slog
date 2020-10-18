package slog

import "testing"

func TestHandler(t *testing.T) {
	h := NewHandlers()
	h.Add(handlerTest)
	err := test(t, []interface{}{stringTest, pointerTest, errorTest, intTest}, h.run(stringTest, pointerTest, errorTest, intTest))
	if err != nil {
		t.Error(err)
	}
}

func handlerTest(v ...interface{}) []interface{} {
	return v
}
