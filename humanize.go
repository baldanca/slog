package main

import (
	"fmt"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

var (
	humanizeFlag bool = false
)

// EnableHumanize function
func EnableHumanize() {
	humanizeFlag = true
}

func humanizeAll(i ...interface{}) []interface{} {
	if !humanizeFlag {
		return i
	}

	for index, data := range i {
		dataH, err := humanize(data)
		if err != nil {
			Warn(err)
			continue
		}
		i[index] = dataH
	}

	return i
}

func humanize(i interface{}) (string, error) {

	// Verifying if is an error type
	if err, ok := i.(error); ok {
		return err.Error(), nil
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
			return "", err
		}

		// Returning JSON data
		return string(bytes), nil

	// Map
	case reflect.Map:
		// Transforming data to JSON
		bytes, err := jsoniter.Marshal(v.Interface())
		if err != nil {
			return "", err
		}

		// Returning JSON data
		return string(bytes), nil

	// Array
	case reflect.Array:
		if av, ok := i.([]byte); ok {
			return string(av), nil
		}

		// Transforming data to JSON
		bytes, err := jsoniter.Marshal(v.Interface())
		if err != nil {
			return "", err
		}

		// Returning JSON data
		return string(bytes), nil

	// Slice
	case reflect.Slice:
		if av, ok := i.([]byte); ok {
			return string(av), nil
		}

		// Transforming data to JSON
		bytes, err := jsoniter.Marshal(v.Interface())
		if err != nil {
			return "", err
		}

		// Returning JSON data
		return string(bytes), nil

	// Default
	default:
		return fmt.Sprintf("%v", i), nil
	}
}
