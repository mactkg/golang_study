package main

import (
	"reflect"
	"unsafe"
)

func CheckCycled(x interface{}) bool {
	seen := make(logMap)
	return check(reflect.ValueOf(x), seen)
}

func check(x reflect.Value, seen logMap) bool {
	if x.CanAddr() {
		ptr := unsafe.Pointer(x.UnsafeAddr())
		c := log{ptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return check(x.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if check(x.Index(i), seen) {
				return true
			}
		}
		return false
	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			if check(x.Field(i), seen) {
				return true
			}
		}
		return false
	case reflect.Map:
		for _, k := range x.MapKeys() {
			if check(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

type log struct {
	x unsafe.Pointer
	t reflect.Type
}

type logMap map[log]bool
