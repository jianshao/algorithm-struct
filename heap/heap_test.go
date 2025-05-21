package heap

import (
	"testing"

	"github.com/jianshao/datastruct/utils/test"
)

func compare(a, b interface{}) int {
	return a.(int) - b.(int)
}

type heapParam struct {
	operate string
	val     interface{}
}

func Test_MinHeap(t *testing.T) {
	heap, _ := Init(10, false, compare)
	params := []test.TestParam{
		{Input: heapParam{operate: "get"}, Expect: nil, ShouldEqual: true},
		{Input: heapParam{operate: "add", val: 8}, Expect: nil, ShouldEqual: true},
		{Input: heapParam{operate: "get"}, Expect: 8, ShouldEqual: true},
		{Input: heapParam{operate: "add", val: 5}, Expect: nil, ShouldEqual: true},
		{Input: heapParam{operate: "add", val: 8}, Expect: nil, ShouldEqual: true},
		{Input: heapParam{operate: "get"}, Expect: 5, ShouldEqual: true},
		{Input: heapParam{operate: "add", val: 7}, Expect: nil, ShouldEqual: true},
		{Input: heapParam{operate: "get"}, Expect: 7, ShouldEqual: true},
	}
	test.Test("min heap", t, params, func(input interface{}) interface{} {
		param := input.(heapParam)
		switch param.operate {
		case "add":
			return heap.Add(param.val)
		case "get":
			result, _ := heap.GetTheTop()
			return result
		}
		return nil
	})
}
