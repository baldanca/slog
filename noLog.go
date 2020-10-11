package slog

import (
	"reflect"
)

// noLog remove data from properties with flag noLog:"true"
func noLogAll(i ...interface{}) []interface{} {
	for index, data := range i {
		if data == nil {
			continue
		}
		i[index] = noLog(data)
	}

	return i
}

func noLog(data interface{}) interface{} {

	// Getting type of data interface
	t := reflect.TypeOf(data)

	// New reflect object based on data type of
	v := reflect.New(t)

	// Verifying if object is a pointer
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		// Getting type of data interface
		t = reflect.TypeOf(data).Elem()

		// New reflect object based on data type of
		v = reflect.New(t)

		// Setting value of reflect object
		v.Elem().Set(reflect.ValueOf(data).Elem())
	} else {
		// Setting value of reflect object
		v.Elem().Set(reflect.ValueOf(data))
	}

	// Verifying kind of object
	switch v.Elem().Kind() {

	// Pointer
	case reflect.Ptr:
		// Getting pointer value
		v = v.Elem()

		// Verifying type of value
		for v.Kind() == reflect.Ptr {
			// if value is a pointer getting value again
			v = v.Elem()
		}

		// Setting new filtred value
		v.Set(reflect.ValueOf(noLog(v.Interface())))
		break

	// Unsafe Pointer
	case reflect.UnsafePointer:
		// Getting pointer value
		v = v.Elem()

		// Verifying type of value
		for v.Kind() == reflect.Ptr {
			// if value is a pointer getting value again
			v = v.Elem()
		}

		// Setting new filtred value
		v.Set(reflect.ValueOf(noLog(v.Interface())))
		break

	// Struct
	case reflect.Struct:
		// Label - https://medium.com/golangspec/labels-in-go-4ffd81932339
	FIELD:
		// Reading all fields
		for i := 0; i < t.NumField(); i++ {

			// Verifying if field can be set
			if !v.Elem().Field(i).CanSet() {
				panic("cannot set field")
			}

			// Getting tag value
			if t.Field(i).Tag.Get("noLog") == "true" {
				// Setting field value to empty
				v.Elem().Field(i).Set(reflect.Zero(v.Elem().Field(i).Type()))

				// Continue to FIELD next loop
				continue FIELD
			}

			// Setting new filtred value
			v.Elem().Field(i).Set(reflect.ValueOf(noLog(v.Elem().Field(i).Interface())))
		}
		break
	case reflect.Array:
		// Reading all fields
		for i := 0; i < v.Elem().Len(); i++ {
			// Setting new filtred value
			v.Elem().Index(i).Set(reflect.ValueOf(noLog(v.Elem().Index(i).Interface())))
		}
		break
	case reflect.Map:
		// Reading map fields
		keys := v.Elem().MapKeys()

		// Reading all fields
		for i := 0; i < v.Elem().Len(); i++ {
			// Setting new filtred value
			v.Elem().SetMapIndex(keys[i], reflect.ValueOf(noLog(v.Elem().MapIndex(keys[i]).Interface())))
		}
		break
	case reflect.Slice:
		// Reading all fields
		for i := 0; i < v.Elem().Len(); i++ {
			// Calling noLog to founded value
			r := noLog(v.Elem().Index(i).Interface())

			if v.Elem().Index(i).Kind() == reflect.Ptr {
				v.Elem().Index(i).Elem().Set(reflect.ValueOf(r).Elem())
				continue
			}

			// Setting new filtred value
			v.Elem().Index(i).Set(reflect.ValueOf(noLog(v.Elem().Index(i).Interface())))
		}
		break
	default:
		break
	}

	// Verifying if object is a pointer
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		// Returning pointer data
		return v.Elem().Interface()
	}

	// Returning data
	return v.Interface()
}
