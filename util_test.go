package slog

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

// 	- TODO
// 		- Review logic

type (
	DataTest struct {
		String  string
		Pointer *string
		Int     int
		Error   error
		JSON    []byte
		Func    func(teste string) bool
		Map     map[string]interface{}
		Array   [7]interface{}
		Slice   []interface{}
	}

	InterTest interface {
		t(in interface{}) (out interface{})
	}
)

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
	structTeste = DataTest{
		String:  stringTest,
		Pointer: pointerTest,
		Int:     intTest,
		Error:   errorTest,
		JSON:    jsonTest,
		Func:    funcTest,
		Map:     mapTest,
		Array:   arrayTest,
		Slice:   sliceTest,
	}
	structTesteWithNil = DataTest{
		String:  "",
		Pointer: nil,
		Int:     0,
		Error:   nil,
		JSON:    nil,
		Func:    nil,
		Map:     nil,
		Array:   [7]interface{}{},
		Slice:   nil,
	}
)

func test(t *testing.T, expectedData, testData interface{}) error {
	var (
		resultData = humanize(testData)
	)

	// Zero
	if expectedData == nil {
		if testData == nil {
			return nil
		}
	}
	if expectedData == nil {
		if reflect.ValueOf(testData).IsZero() {
			return nil
		}
	}
	if reflect.ValueOf(expectedData).IsZero() {
		if testData == nil {
			return nil
		}
	}
	if reflect.ValueOf(expectedData).IsZero() {
		if reflect.ValueOf(testData).IsZero() {
			return nil
		}
	}

	// Error
	if err, ok := expectedData.(error); ok {
		if err.Error() != fmt.Sprintf("%v", resultData) {
			t.Log(err, resultData, testData)
			return errors.New("error not match")
		}
		return nil
	}

	switch reflect.ValueOf(expectedData).Type().Kind() {

	// Pointer
	case reflect.Ptr:
		err := test(t, reflect.ValueOf(expectedData).Elem().Interface(), testData)
		if err != nil {
			t.Error(err)
		}
		break

	// Struct
	case reflect.Struct:
		resultObject := new(map[string]interface{})
		err := jsoniter.Unmarshal([]byte(resultData), &resultObject)
		if err != nil {
			return err
		}
		for numField := 0; numField < reflect.ValueOf(expectedData).NumField(); numField++ {
			nameField := reflect.TypeOf(expectedData).Field(numField).Name
			ve := reflect.ValueOf(expectedData).FieldByName(nameField).Interface()
			vr := reflect.ValueOf(resultObject).Elem().MapIndex(reflect.ValueOf(nameField)).Elem().Interface()
			err := test(t, ve, vr)
			if err != nil {
				return err
			}
		}
		return nil

	// Map
	case reflect.Map:
		resultObject := make(map[string]interface{})
		oks := reflect.ValueOf(expectedData).Len()
		err := jsoniter.Unmarshal([]byte(resultData), &resultObject)
		if err != nil {
			return err
		}
		for ke, ve := range expectedData.(map[string]interface{}) {
			for kr, vr := range resultObject {
				if ke == kr {
					err := test(t, ve, vr)
					if err != nil {
						return err
					}
					oks--
				}
			}
		}
		if oks != 0 {
			return fmt.Errorf("%s %d", "map not match", oks)
		}
		break

	// Array
	case reflect.Array:
		resultObject := make([]interface{}, reflect.ValueOf(expectedData).Len())
		oks := reflect.ValueOf(expectedData).Len()
		err := jsoniter.Unmarshal([]byte(resultData), &resultObject)
		if err != nil {
			return err
		}
		for ke := 0; ke < len(resultObject); ke++ {
			for kr, vr := range resultObject {
				if ke == kr {
					ve := reflect.ValueOf(expectedData).Index(ke).Elem().Interface()
					err := test(t, ve, vr)
					if err != nil {
						return err
					}
					oks--
				}
			}
		}
		if oks != 0 {
			return fmt.Errorf("%s %d", "array not match", oks)
		}
		break

	// Slice
	case reflect.Slice:
		if jsonv, ok := expectedData.([]byte); jsoniter.Valid(jsonv) && ok {
			expectedObject := new(interface{})
			err := jsoniter.Unmarshal(jsonv, &expectedObject)
			if err != nil {
				return err
			}
			err = test(t, expectedObject, resultData)
			if err != nil {
				return err
			}
			return nil
		}
		resultObject := []interface{}{}
		oks := reflect.ValueOf(expectedData).Len()
		err := jsoniter.Unmarshal([]byte(resultData), &resultObject)
		if err != nil {
			return err
		}
		for ke, ve := range expectedData.([]interface{}) {
			for kr, vr := range resultObject {
				if ke == kr {
					err := test(t, ve, vr)
					if err != nil {
						return err
					}
					oks--
				}
			}
		}
		if oks != 0 {
			return fmt.Errorf("%s %d", "map not match", oks)
		}
		break
	case reflect.Func:
		if fmt.Sprintf("%v", resultData) == "" {
			return nil
		}

	// Default
	default:
		if fmt.Sprintf("%v", expectedData) != fmt.Sprintf("%v", resultData) {
			t.Log(fmt.Sprintf("%v", expectedData), reflect.TypeOf(expectedData), fmt.Sprintf("%v", resultData), reflect.TypeOf(resultData))
			return errors.New("data not match")
		}
	}
	return nil
}
