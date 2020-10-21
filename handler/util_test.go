package handler

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/fatih/structs"
	jsoniter "github.com/json-iterator/go"
)

// 	- TODO
// 		- Review logic

type (
	dataTest struct {
		String  string                  `nolog:"true"`
		Pointer *string                 `nolog:"true"`
		Int     int                     `nolog:"true"`
		Error   error                   `nolog:"true"`
		JSON    []byte                  `nolog:"true"`
		Func    func(teste string) bool `nolog:"true"`
		Map     map[string]interface{}  `nolog:"true"`
		Array   [7]interface{}          `nolog:"true"`
		Slice   []interface{}           `nolog:"true"`
		Logged  string
	}
	interTest interface {
		t(in interface{}) (out interface{})
	}
)

func test(t *testing.T, expectedData, testData interface{}) {
	if expectedData == nil && testData == nil {
		return
	}
	expectedElem := reflect.ValueOf(expectedData)
	if expectedElem.IsZero() && testData == nil {
		return
	}
	testElem := reflect.ValueOf(testData)
	if expectedData == nil && testElem.IsZero() {
		return
	}
	if expectedElem.IsZero() && testElem.IsZero() {
		return
	}
	if err, ok := expectedData.(error); ok {
		if err.Error() != fmt.Sprintf("%v", testData) {
			t.Log("error:", err, "test:", testData)
			t.Error(errors.New("error not match"))
			return
		}
		return
	}
	switch expectedElem.Kind() {
	case reflect.Array:
		expectedArrayLength := expectedElem.Len()
		testArrayLength := testElem.Len()
		if expectedArrayLength != testArrayLength {
			t.Log(expectedData, testData)
			t.Error("length is not equal")
			return
		}
		testOK := expectedArrayLength
		for expectedIndex := 0; expectedIndex < expectedArrayLength; expectedIndex++ {
			for testIndex := 0; testIndex < testArrayLength; testIndex++ {
				if expectedIndex == testIndex {
					expectedValue := expectedElem.Index(expectedIndex).Elem().Interface()
					testValue := testElem.Index(testIndex).Elem().Interface()
					test(t, expectedValue, testValue)
					testOK--
				}
			}
		}
		if testOK != 0 {
			t.Log(expectedData, testData)
			t.Error("testOK != 0")
			return
		}
		return
	case reflect.Slice:
		var (
			expectedObject interface{}
			testObject     interface{}
			jsonOK         bool
		)
		if expectedBytes, ok := expectedData.([]byte); jsoniter.Valid(expectedBytes) && ok {
			err := jsoniter.Unmarshal(expectedBytes, &expectedObject)
			if err != nil {
				t.Error(err)
				return
			}
			jsonOK = true
			return
		}
		if testBytes, ok := testData.([]byte); jsoniter.Valid(testBytes) && ok {
			err := jsoniter.Unmarshal(testBytes, &testObject)
			if err != nil {
				t.Error(err)
				return
			}
			jsonOK = true
			return
		}
		if expectedObject == nil {
			expectedObject = expectedData
		}
		if testObject == nil {
			testObject = testData
		}
		if jsonOK {
			test(t, expectedObject, testObject)
		}
		expectedArrayLength := expectedElem.Len()
		testArrayLength := testElem.Len()
		if expectedArrayLength != testArrayLength {
			t.Log(expectedData, testData)
			t.Error("length is not equal")
			return
		}
		testOK := expectedArrayLength
		for expectedIndex := 0; expectedIndex < expectedArrayLength; expectedIndex++ {
			for testIndex := 0; testIndex < testArrayLength; testIndex++ {
				if expectedIndex == testIndex {
					expectedValue := expectedElem.Index(expectedIndex).Elem().Interface()
					testValue := testElem.Index(testIndex).Elem().Interface()
					test(t, expectedValue, testValue)
					testOK--
				}
			}
		}
		if testOK != 0 {
			t.Log(expectedData, testData)
			t.Error("testOK != 0")
			return
		}
		return
	case reflect.Ptr:
		expectedElem = expectedElem.Elem()
		test(t, expectedElem.Interface(), testElem.Interface())
		return
	case reflect.Struct:
		switch testElem.Kind() {
		case reflect.Ptr:
			testElem = testElem.Elem()
			test(t, expectedElem.Interface(), testElem.Interface())
			return
		case reflect.Struct:
			expectedNumFields := expectedElem.NumField()
			testNumFields := testElem.NumField()
			testOK := expectedNumFields
			if expectedNumFields != testNumFields {
				t.Log(expectedData, testData)
				t.Error("numFields is not equal")
				return
			}
			for expectedNumField := 0; expectedNumField < expectedNumFields; expectedNumField++ {
				for testNumField := 0; testNumField < testNumFields; testNumField++ {
					expectedNameField := expectedElem.Type().Field(expectedNumField).Name
					testNameField := testElem.Type().Field(testNumField).Name
					if expectedNameField == testNameField {
						expectedValueField := expectedElem.FieldByName(expectedNameField).Interface()
						testValueField := testElem.FieldByName(testNameField).Interface()
						test(t, expectedValueField, testValueField)
						testOK--
					}
				}
			}
			if testOK != 0 {
				t.Log(expectedData, testData)
				t.Error("testOK != 0")
				return
			}
			return
		case reflect.Map:
			expectedMap := structs.Map(expectedElem.Interface())
			expectedMapElem := reflect.ValueOf(expectedMap)
			expectedMapLength := expectedMapElem.Len()
			testMapLength := testElem.Len()
			if expectedMapLength != testMapLength {
				t.Log(expectedData, testData)
				t.Error("numFields is not equal")
				return
			}
			testOK := expectedMapLength
			expectedMapIter := expectedMapElem.MapRange()
			for {
				if !expectedMapIter.Next() {
					break
				}
				expectedMapKey := expectedMapIter.Key()
				expectedMapValue := expectedMapIter.Value()
				testMapValue := testElem.MapIndex(expectedMapKey)
				test(t, expectedMapValue, testMapValue)
				testOK--
			}
			if testOK != 0 {
				t.Log(expectedData, testData)
				t.Error("testOK != 0")
				return
			}
			return
		default:
			t.Error("data is not same type")
			return
		}
	case reflect.Map:
		expectedMapLength := expectedElem.Len()
		testMapLength := testElem.Len()
		if expectedMapLength != testMapLength {
			t.Log(expectedData, testData)
			t.Error("numFields is not equal")
			return
		}
		testOK := expectedMapLength
		expectedMapIter := expectedElem.MapRange()
		for {
			if !expectedMapIter.Next() {
				break
			}
			expectedMapKey := expectedMapIter.Key()
			expectedMapValue := expectedMapIter.Value()
			testMapValue := testElem.MapIndex(expectedMapKey)
			test(t, expectedMapValue.Interface(), testMapValue.Interface())
			testOK--
		}
		if testOK != 0 {
			t.Log(expectedData, testData)
			t.Error("testOK != 0")
			return
		}
		return
	default:
		switch testElem.Kind() {
		case reflect.Ptr:
			testElem = testElem.Elem()
			test(t, expectedElem.Interface(), testElem.Interface())
			return
		default:
			if fmt.Sprintf("%v", expectedData) != fmt.Sprintf("%v", testData) {
				t.Log(fmt.Sprintf("%v", expectedData), expectedElem.Type(), fmt.Sprintf("%v", testData), testElem.Type())
				t.Error(errors.New("data not match"))
				return
			}
		}
	}
}
