package search

import (
	"testing"

	"github.com/jianshao/datastruct/utils/test"
)

type binSearchParams struct {
	arr []interface{}
	val interface{}
}

func intLess(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func stringLess(a, b interface{}) bool {
	return a.(string) < b.(string)
}

func Test_BinSearch_Int(t *testing.T) {
	params := []test.TestParam{
		{Input: binSearchParams{arr: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, val: 6}, Expect: 5, ShouldEqual: true},
		{Input: binSearchParams{arr: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, val: 0}, Expect: -1, ShouldEqual: true},
	}
	test.Test("测试整型数BinSearch", t, params, func(input interface{}) interface{} {
		param := input.(binSearchParams)
		return BinSearch(param.arr, param.val, intLess)
	})
}

func Test_BinSearch_String(t *testing.T) {
	params := []test.TestParam{
		{
			Input:       binSearchParams{arr: []interface{}{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, val: "6"},
			Expect:      5,
			ShouldEqual: true,
		},
		{
			Input:       binSearchParams{arr: []interface{}{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, val: "0"},
			Expect:      -1,
			ShouldEqual: true,
		},
		{
			Input:       binSearchParams{arr: []interface{}{"1", "12", "123", "1234", "2", "22", "7", "8", "9"}, val: "12"},
			Expect:      1,
			ShouldEqual: true,
		},
	}
	test.Test("测试整型数BinSearch", t, params, func(input interface{}) interface{} {
		param := input.(binSearchParams)
		return BinSearch(param.arr, param.val, stringLess)
	})
}
