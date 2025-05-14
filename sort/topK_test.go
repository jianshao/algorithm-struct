package sort

import (
	"testing"

	"github.com/jianshao/datastruct/utils/test"
)

type TestTopKParam struct {
	arr []int
	kth int
}

func Test_TopK(t *testing.T) {
	params := []test.TestParam{
		{Input: TestTopKParam{arr: []int{1, 3, 6, 3, 8, 22, 545}, kth: 2}, Expect: 3, ShouldEqual: true},
		{Input: TestTopKParam{arr: []int{1, 3, 6, 3, 8, 22, 545}, kth: 2}, Expect: 3, ShouldEqual: false},
	}
	test.Test("topk", t, params, func(input interface{}) interface{} {
		info := input.(TestTopKParam)
		return TopK(info.arr, info.kth)
	})
}
