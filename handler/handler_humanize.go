package handler

import (
	"fmt"
	"reflect"

	"github.com/fatih/structs"
	jsoniter "github.com/json-iterator/go"
)

// HumanizeAll informations
func HumanizeAll(values ...interface{}) []interface{} {
	for index, data := range values {
		values[index] = humanize(data)
	}
	return values
}

func humanize(value interface{}) string {
	if value == nil {
		return ""
	}
	if err, ok := value.(error); ok {
		return err.Error()
	}
	valueType := reflect.TypeOf(value)
	valueElem := reflect.New(valueType).Elem()
	valueElem.Set(reflect.ValueOf(value))
	switch valueElem.Kind() {
	case reflect.Array:
		for index := 0; index < valueElem.Len(); index++ {
			indexValue := valueElem.Index(index)
			humanizedValue := humanize(indexValue.Interface())
			indexValue.Set(reflect.ValueOf(humanizedValue))
		}
		byteValue, _ := jsoniter.Marshal(valueElem.Interface())
		return string(byteValue)
	case reflect.Slice:
		if valueElem.IsNil() {
			return ""
		}
		if byteValue, ok := value.([]byte); jsoniter.Valid(byteValue) && ok {
			return string(byteValue)
		}
		for index := 0; index < valueElem.Len(); index++ {
			indexValue := valueElem.Index(index)
			humanizedValue := humanize(indexValue.Interface())
			indexValue.Set(reflect.ValueOf(humanizedValue))
		}
		byteValue, _ := jsoniter.Marshal(valueElem.Interface())
		return string(byteValue)
	case reflect.Ptr:
		if valueElem.IsZero() {
			return ""
		}
		return humanize(valueElem.Elem().Interface())
	case reflect.Struct:
		mapValue := structs.Map(valueElem.Interface())
		mapValueElem := reflect.ValueOf(mapValue)
		for _, mapKey := range mapValueElem.MapKeys() {
			mapKeyValue := mapValueElem.MapIndex(mapKey).Interface()
			humanizedValue := humanize(mapKeyValue)
			mapValueElem.SetMapIndex(mapKey, reflect.ValueOf(humanizedValue))
		}
		byteValue, _ := jsoniter.Marshal(mapValue)
		return string(byteValue)
	case reflect.Map:
		if valueElem.IsNil() {
			return ""
		}
		for _, mapKey := range valueElem.MapKeys() {
			mapKeyValue := valueElem.MapIndex(mapKey).Interface()
			humanizedValue := humanize(mapKeyValue)
			valueElem.SetMapIndex(mapKey, reflect.ValueOf(humanizedValue))
		}
		byteValue, _ := jsoniter.Marshal(valueElem.Interface())
		return string(byteValue)
	case reflect.Func:
		if valueElem.IsNil() {
			return ""
		}
		// Can't do better than this:
		return ""
	default:
		return fmt.Sprintf("%v", value)
	}
}
