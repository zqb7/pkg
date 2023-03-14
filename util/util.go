package util

import "reflect"

// 非指针类型的interface,转为其对应的指针类型interface
func InterfaceTypePtr(obj interface{}) interface{} {
	val := reflect.ValueOf(obj)
	vp := reflect.New(val.Type())
	vp.Elem().Set(val)
	return vp.Interface()
}
