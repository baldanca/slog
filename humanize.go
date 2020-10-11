package slog

import (
	"fmt"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

// humanizeAll informations of an interface array
func humanizeAll(i ...interface{}) []interface{} {
	for index, data := range i {
		if data == nil {
			continue
		}
		i[index] = humanize(data)
	}

	return i
}

// humanize information
func humanize(i interface{}) string {

	// Verifying if is an error type
	if err, ok := i.(error); ok {
		return err.Error()
	}

	// Getting type of data interface
	t := reflect.TypeOf(i)

	// New reflect object based on data type of
	v := reflect.New(t)

	// Setting value of reflect object
	v.Elem().Set(reflect.ValueOf(i))

	// Verifying kind of object
	switch v.Elem().Kind() {

	// Pointer
	case reflect.Ptr:
		// Verifying type of value
		for v.Kind() == reflect.Ptr {
			// if value is a pointer getting value again
			v = v.Elem()
		}

		// Calling noLog value
		return humanize(v.Interface())

	// Struct
	case reflect.Struct:
		// Transforming data to JSON
		bytes, err := jsoniter.Marshal(v.Interface())
		if err != nil {
			panic(err)
		}

		// Returning JSON data
		return string(bytes)

	// Map
	case reflect.Map:
		// Transforming data to JSON
		bytes, err := jsoniter.Marshal(v.Interface())
		if err != nil {
			panic(err)
		}

		// Returning JSON data
		return string(bytes)

	// Array
	case reflect.Array:
		if av, ok := i.([]byte); ok {
			return string(av)
		}

		// Transforming data to JSON
		bytes, err := jsoniter.Marshal(v.Interface())
		if err != nil {
			panic(err)
		}

		// Returning JSON data
		return string(bytes)

	// Slice
	case reflect.Slice:
		if av, ok := i.([]byte); ok {
			return string(av)
		}

		// Transforming data to JSON
		bytes, err := jsoniter.Marshal(v.Interface())
		if err != nil {
			panic(err)
		}

		// Returning JSON data
		return string(bytes)

	// Default
	default:
		return fmt.Sprintf("%v", i)
	}
}
