package utils

import "reflect"

type Compare func(a, b interface{}) bool

func CompareArr(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if reflect.TypeOf(a[i]) != reflect.TypeOf(b[i]) {
			return false
		}
		if reflect.ValueOf(a[i]) != reflect.ValueOf(b[i]) {
			return false
		}
	}

	return true
}

func Arr2Interfaces(arr []int) []interface{} {
	result := make([]interface{}, 0)
	for i := 0; i < len(arr); i++ {
		result = append(result, arr[i])
	}
	return result
}

func CompareIntArr(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
