package slog

import (
	"fmt"
	"reflect"

	"github.com/fatih/structs"
	jsoniter "github.com/json-iterator/go"
)

// humanizeAll informations of an interface array
func humanizeAll(i ...interface{}) []interface{} {
	for index, data := range i {
		i[index] = humanize(data)
	}
	return i
}

// humanize information
func humanize(i interface{}) string {
	// Verifying nil value:
	if i == nil {
		return ""
	}
	// Verifying if is an error type:
	if err, ok := i.(error); ok {
		return err.Error()
	}
	// Getting type of data interface:
	t := reflect.TypeOf(i)
	// New reflect object based on data type of:
	v := reflect.New(t).Elem()
	// Setting value of data interface:
	v.Set(reflect.ValueOf(i))
	// Verifying kind of object:
	switch v.Kind() {
	case reflect.Array:
		// Humanizing array elements:
		for i := 0; i < v.Len(); i++ {
			val := v.Index(i)
			hVal := humanize(val.Interface())
			val.Set(reflect.ValueOf(hVal))
		}
		// Transforming data to JSON:
		bytes, _ := jsoniter.Marshal(v.Interface())
		// Returning JSON data:
		return string(bytes)
	case reflect.Slice:
		// Verifying if slice is nil:
		if v.IsNil() {
			return ""
		}
		// Verifying json type:
		if bytes, ok := i.([]byte); jsoniter.Valid(bytes) && ok {
			return string(bytes)
		}
		// Humanizing array elements:
		for i := 0; i < v.Len(); i++ {
			val := v.Index(i)
			hVal := humanize(val.Interface())
			val.Set(reflect.ValueOf(hVal))
		}
		// Transforming data to JSON:
		bytes, _ := jsoniter.Marshal(v.Interface())
		// Returning JSON data:
		return string(bytes)
	case reflect.Ptr:
		// Verifying if pointer is zero:
		if v.IsZero() {
			return ""
		}
		// Calling humanize again:
		return humanize(v.Elem().Interface())
	case reflect.Struct:
		// Transforming struct to map[string]interface{}:
		sMap := structs.Map(v.Interface())
		vMap := reflect.ValueOf(sMap)
		// Humanizing map elements:
		for _, key := range vMap.MapKeys() {
			val := vMap.MapIndex(key).Interface()
			hVal := humanize(val)
			vMap.SetMapIndex(key, reflect.ValueOf(hVal))
		}
		// Transforming data to JSON:
		bytes, _ := jsoniter.Marshal(sMap)
		// Returning JSON data:
		return string(bytes)
	case reflect.Map:
		// Verifying if slice is nil:
		if v.IsNil() {
			return ""
		}
		// Humanizing map elements:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key).Interface()
			hVal := humanize(val)
			v.SetMapIndex(key, reflect.ValueOf(hVal))
		}
		// Transforming data to JSON:
		bytes, _ := jsoniter.Marshal(v.Interface())
		// Returning JSON data:
		return string(bytes)
	case reflect.Func:
		if v.IsNil() {
			return ""
		}
		// Can't do better than this:
		return ""
	default:
		return fmt.Sprintf("%v", i)
	}
}
