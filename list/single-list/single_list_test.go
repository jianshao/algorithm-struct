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

func Test_SingleList(t *testing.T) {
	// l := Init(compare)
	var l *SingleList
	params := []test.TestParam{
		// single list basic tests
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

		// merge parter tests

		// reverse parter tests
		{Input: singleListParam{operate: "init"}, Expect: 1, ShouldEqual: true},
		// only 1 node reverse
		{Input: singleListParam{operate: "add", val: 1}, Expect: nil, ShouldEqual: true},
		{Input: singleListParam{operate: "reverse"}, Expect: nil, ShouldEqual: true},
		{Input: singleListParam{operate: "get", val: 1}, Expect: 1, ShouldEqual: true},

		// x nodes reverse
		{Input: singleListParam{operate: "add", val: 6}, Expect: nil, ShouldEqual: true},
		{Input: singleListParam{operate: "reverse"}, Expect: nil, ShouldEqual: true},
		{Input: singleListParam{operate: "get", val: 1}, Expect: 6, ShouldEqual: true},
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
			result, _ := l.IsExist(param.val)
			return result
		case "get":
			result, _ := l.GetTheNthNode(param.val.(int))
			return result
		case "count":
			result, _ := l.GetNodesCount()
			return result
		case "reverse":
			l.Reverse()
			return nil
		default:
			return nil
		}
	})
}
