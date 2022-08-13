package main

import "reflect"

func main() {}

func through(x interface{}, fn func(input string)) {
	value := getValue(x)

	throughValue := func(value reflect.Value) {
		through(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			throughValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			throughValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			throughValue(value.MapIndex(key))
		}
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
