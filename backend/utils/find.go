package utils

import "reflect"

func Find[T comparable](slice []T, elem T) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

func FindFunc[T any](slice []T, f func(T) bool) (T, bool) {
	for _, v := range slice {
		if f(v) {
			return v, true
		}
	}
	return reflect.Zero(reflect.TypeOf(slice).Elem()).Interface().(T), false
}
