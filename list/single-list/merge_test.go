package singlelist

import (
	"testing"

	"github.com/jianshao/datastruct/utils/test"
)

type singleListParam struct {
	operate string
	val     interface{}
}

func compare(a, b interface{}) int {
	return a.(int) - b.(int)
}

func Test_Merge(t *testing.T) {
	// l := Init(compare)
	var l *SingleListHead
	params := []test.TestParam{
		{Input: singleListParam{operate: "add", val: 1}, Expect: nil, ShouldEqual: false},
		{Input: singleListParam{operate: "init", val: 1}, Expect: 1, ShouldEqual: true},
		{Input: singleListParam{operate: "check", val: 1}, Expect: false, ShouldEqual: true},
		{Input: singleListParam{operate: "add", val: 1}, Expect: nil, ShouldEqual: true},
		{Input: singleListParam{operate: "check", val: 1}, Expect: true, ShouldEqual: true},
		{Input: singleListParam{operate: "add", val: 6}, Expect: nil, ShouldEqual: true},
		{Input: singleListParam{operate: "count"}, Expect: 2, ShouldEqual: true},
		{Input: singleListParam{operate: "check", val: 6}, Expect: true, ShouldEqual: true},
		{Input: singleListParam{operate: "remove", val: 1}, Expect: 1, ShouldEqual: true},
		{Input: singleListParam{operate: "count"}, Expect: 1, ShouldEqual: true},
		{Input: singleListParam{operate: "check", val: 1}, Expect: false, ShouldEqual: true},
	}
	test.Test("single list merge", t, params, func(input interface{}) interface{} {
		param := input.(singleListParam)
		switch param.operate {
		case "init":
			l = Init(compare)
			return 1
		case "add":
			return l.Append(param.val)
		case "remove":
			l.Remove(param.val)
			return 1
		case "check":
			return l.IsExist(param.val)
		case "count":
			return l.GetNodesCount()
		default:
			return nil
		}
	})
}
