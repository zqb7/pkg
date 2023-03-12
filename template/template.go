package template

import (
	"reflect"
)

func Gen(t any) (result any) {
	rt := reflect.TypeOf(t)
	switch rt.Kind() {
	case reflect.Struct:
		result = newValue(t, reflect.ValueOf(rt)).Interface()
	case reflect.Slice:
		rSlice := reflect.MakeSlice(rt, 1, 1)
		rSlice.Index(0).Set(newValue(rSlice.Index(0).Interface(), rSlice.Index(0)))
		result = rSlice.Interface()
	case reflect.Pointer:
		result = newValue(t, reflect.ValueOf(&t).Elem()).Interface()
	}
	return result
}

func newValue(obj interface{}, rv reflect.Value) reflect.Value {
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Pointer:
		switch rt.Elem().Kind() {
		case reflect.Struct:
			return newValue(reflect.Indirect(reflect.New(rt.Elem())).Interface(), reflect.ValueOf(rv))
		default:
			return reflect.New(rt.Elem())
		}
	case reflect.Struct:
		newObj := reflect.New(rt).Interface()
		newRV := reflect.ValueOf(newObj)
		elem := newRV.Elem()
		for index := 0; index < elem.NumField(); index++ {
			field := elem.Field(index)
			field.Set(newValue(field.Interface(), field))
		}
		return elem
	case reflect.Interface:
		return newValue(obj, reflect.ValueOf(rv))
	}
	return rv
}
