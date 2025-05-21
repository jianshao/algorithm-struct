package heap

import (
	"testing"

	"github.com/jianshao/datastruct/utils/test"
)

type heapSortParam struct {
	arr []int
}

func Test_Heap_Sort(t *testing.T) {
	params := []test.TestParam{
		{Input: heapSortParam{arr: []int{9, 4, 6, 2, 3, 1, 5, 8, 7}}, Expect: [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, ShouldEqual: true},
	}
	test.Test("heap sort", t, params, func(input interface{}) interface{} {
		param := input.(heapSortParam)
		err := HeapSort(param.arr)
		if err != nil {
			return false
		}
		// reflect中slice不能比较，需要转换成array。
		result := [9]int{}
		// slice中元素个数不确定，可以这样写，安全点。
		copy(result[:], param.arr[0:9])
		return result
	})
}
