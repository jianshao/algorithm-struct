package search

import "reflect"

type Less func(a, b interface{}) bool

// 前提是数组是有序的，返回的数据在数组中的下标，-1表示没找到
func BinSearch(arr []interface{}, val interface{}, less Less) int {
	left, right := 0, len(arr)-1
	for left < right {
		if reflect.ValueOf(arr[left]).Equal(reflect.ValueOf(val)) {
			return left
		}
		if reflect.ValueOf(arr[right]).Equal(reflect.ValueOf(val)) {
			return right
		}
		mid := (left + right) / 2
		if reflect.ValueOf(arr[mid]).Equal(reflect.ValueOf(val)) {
			return mid
		}
		if less(arr[mid], val) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
