package sort

import (
	"testing"

	"github.com/jianshao/datastruct/utils"
)

func intLess(a, b interface{}) bool {
	valA := a.(int)
	valB := b.(int)
	if valA < valB {
		return true
	} else {
		return false
	}
}

func stringLess(a, b interface{}) bool {
	strA := a.(string)
	strB := b.(string)
	if strA < strB {
		return true
	} else {
		return false
	}
}

func Test_Qsort_Int(t *testing.T) {
	result := utils.Arr2InterfaceArr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	input := utils.Arr2InterfaceArr([]int{9, 4, 5, 8, 2, 1, 3, 7, 6})
	QSort(input, intLess)
	if utils.CompareArr(input, result) {
		t.Errorf("Error: Expect: %v\nGot: %v", result, input)
	}
}

func Test_Qsort_String(t *testing.T) {
	result := utils.Arr2InterfaceArr([]string{"123", "a1", "abc", "abc", "abcd"})
	input := utils.Arr2InterfaceArr([]string{"123", "abcd", "abc", "a1", "abc"})
	QSort(input, stringLess)
	if utils.CompareArr(input, result) {
		t.Errorf("Error: Expect: %v\nGot: %v", result, input)
	}
}
