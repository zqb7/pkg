package template

import (
	"reflect"
)

func Gen(t any) (result any) {
	return newValue(t).Interface()
}

func newValue(obj any) reflect.Value {
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Pointer:
		switch rt.Elem().Kind() {
		case reflect.Struct:
			return newValue(reflect.Indirect(reflect.New(rt.Elem())).Interface())
		default:
			return reflect.New(rt.Elem())
		}
	case reflect.Struct:
		newObj := reflect.New(rt).Interface()
		newRV := reflect.ValueOf(newObj)
		elem := newRV.Elem()
		for index := 0; index < elem.NumField(); index++ {
			if field := elem.Field(index); field.CanSet() {
				fieldSet(field, newValue(field.Interface()))
			}
		}
		return elem
	case reflect.Interface:
		return newValue(obj)
	case reflect.Slice:
		rSlice := reflect.MakeSlice(rt, 1, 1)
		rSlice.Index(0).Set(newValue(rSlice.Index(0).Interface()))
		return rSlice
	}
	return reflect.Zero(rt)
}

func fieldSet(field, value reflect.Value) {
	fieldKind := field.Type().Kind()
	valueKind := value.Type().Kind()
	switch {
	case fieldKind == reflect.Pointer && valueKind == reflect.Struct:
		field.Set(value.Addr())
	default:
		field.Set(value)
	}
}
