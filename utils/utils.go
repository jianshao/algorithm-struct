package utils

import (
	"reflect"
)

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

// 将数组转换为interface数组。
func Arr2InterfaceArr(arr interface{}) []interface{} {
	vals := reflect.ValueOf(arr)
	if vals.Kind() != reflect.Array && vals.Kind() != reflect.Slice {
		// fmt.Printf("type error: %s", vals.Kind().String())
		return nil
	}

	arrLength := vals.Len()
	result := make([]interface{}, arrLength)
	// 数组类型可以作为interface类型，想要转换成interface数组必须逐个元素赋值
	for i := 0; i < arrLength; i++ {
		// 必须要使用最后的interface()，Index()返回的类型是reflect.Value，不是传入的数据类型，业务执行操作的时候会panic
		result[i] = vals.Index(i).Interface()
	}
	return result
}
