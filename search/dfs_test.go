package search

import (
	"testing"

	"github.com/jianshao/datastruct/utils/test"
)

type dfsParam struct {
	board      [][]byte
	startPoint []int
	endPoint   []int
}

func Test_Dfs(t *testing.T) {
	params := []test.TestParam{
		// invalid board
		{Input: dfsParam{board: [][]byte{{1, 1}, {1, 1, 1}, {1, 1, 1}}, startPoint: []int{1, 1}, endPoint: []int{1, 1}}, Expect: -1, ShouldEqual: true},
		{Input: dfsParam{board: [][]byte{{1, 1, 1}, {1, 1}, {1, 1, 1}}, startPoint: []int{1, 1}, endPoint: []int{1, 1}}, Expect: -1, ShouldEqual: true},
		{Input: dfsParam{board: [][]byte{{1, 1, 1}, {1, 1, 1}, {1, 1}}, startPoint: []int{1, 1}, endPoint: []int{1, 1}}, Expect: -1, ShouldEqual: true},
		// invalid point params
		{Input: dfsParam{board: [][]byte{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, startPoint: []int{1, 3}, endPoint: []int{1, 1}}, Expect: -1, ShouldEqual: true},
		{Input: dfsParam{board: [][]byte{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, startPoint: []int{3, 3}, endPoint: []int{1, 1}}, Expect: -1, ShouldEqual: true},
		{Input: dfsParam{board: [][]byte{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, startPoint: []int{1, 0}, endPoint: []int{-1, 1}}, Expect: -1, ShouldEqual: true},
		{Input: dfsParam{board: [][]byte{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, startPoint: []int{1, 0}, endPoint: []int{1, 3}}, Expect: -1, ShouldEqual: true},

		// unreachable end point
		{Input: dfsParam{board: [][]byte{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}}, startPoint: []int{0, 0}, endPoint: []int{2, 2}}, Expect: -1, ShouldEqual: true},
		{Input: dfsParam{board: [][]byte{{1, 1, 1}, {1, 0, 0}, {1, 0, 1}}, startPoint: []int{1, 0}, endPoint: []int{1, 3}}, Expect: -1, ShouldEqual: true},

		// normal tests
		{Input: dfsParam{board: [][]byte{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, startPoint: []int{0, 0}, endPoint: []int{2, 2}}, Expect: 4, ShouldEqual: true},
		{Input: dfsParam{board: [][]byte{{1, 1, 1, 0}, {0, 0, 1, 1}, {1, 1, 1, 0}, {1, 0, 0, 1}, {1, 1, 1, 1}}, startPoint: []int{0, 0}, endPoint: []int{4, 3}}, Expect: 11, ShouldEqual: true},
	}
	test.Test("测试dfs", t, params, func(input interface{}) interface{} {
		param := input.(dfsParam)
		return Dfs(param.board, param.startPoint, param.endPoint)
	})
}
