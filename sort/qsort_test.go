package sort

import (
	"testing"

	"github.com/jianshao/datastruct/utils"
)

func less(a, b interface{}) bool {
	valA := a.(int)
	valB := b.(int)
	if valA < valB {
		return true
	} else {
		return false
	}
}

func Test_Qsort(t *testing.T) {
	result := utils.Arr2Interfaces([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	input := utils.Arr2Interfaces([]int{9, 4, 5, 8, 2, 1, 3, 7, 6})
	QSort(input, less)
	if utils.CompareArr(input, result) {
		t.Errorf("Error: Expect: %v\nGot: %v", result, input)
	}
}
