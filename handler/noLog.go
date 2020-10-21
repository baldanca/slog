package handler

import (
	"reflect"
)

// NoLogAll informations
func NoLogAll(values ...interface{}) []interface{} {
	for index, data := range values {
		values[index] = noLog(data)
	}
	return values
}

func noLog(value interface{}) interface{} {
	if value == nil {
		return value
	}
	valueType := reflect.TypeOf(value)
	valueElem := reflect.New(valueType).Elem()
	valueElem.Set(reflect.ValueOf(value))
	if err, ok := value.(error); ok {
		return err
	}
	switch valueElem.Kind() {
	case reflect.Array:
		for index := 0; index < valueElem.Len(); index++ {
			indexValue := valueElem.Index(index)
			noLogedValue := noLog(indexValue.Interface())
			indexValue.Set(reflect.ValueOf(noLogedValue))
		}
		return valueElem.Interface()
	case reflect.Slice:
		for index := 0; index < valueElem.Len(); index++ {
			indexValue := valueElem.Index(index)
			noLogedValue := noLog(indexValue.Interface())
			indexValue.Set(reflect.ValueOf(noLogedValue))
		}
		return valueElem.Interface()
	case reflect.Ptr:
		return noLog(valueElem.Elem().Interface())
	case reflect.Struct:
		valueNumFields := valueElem.NumField()
		for valueNumField := 0; valueNumField < valueNumFields; valueNumField++ {
			valueField := valueElem.Field(valueNumField)
			noLogTag := valueType.Field(valueNumField).Tag.Get("nolog")
			if noLogTag == "true" {
				zeroValue := reflect.Zero(valueField.Type())
				valueField.Set(zeroValue)
				continue
			}
			noLogedValue := noLog(valueField.Interface())
			valueField.Set(reflect.ValueOf(noLogedValue))
		}
		return valueElem.Interface()
	case reflect.Map:
		for _, mapKey := range valueElem.MapKeys() {
			mapKeyValue := valueElem.MapIndex(mapKey).Interface()
			noLogedValue := noLog(mapKeyValue)
			valueElem.SetMapIndex(mapKey, reflect.ValueOf(noLogedValue))
		}
		return valueElem.Interface()
	default:
		return value
	}
}
