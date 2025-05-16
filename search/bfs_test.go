package search

import (
	"testing"

	"github.com/jianshao/datastruct/utils/test"
)

type bfsParam struct {
	board  [][]int
	endPos []int
}

func Test_Bfs(t *testing.T) {
	params := []test.TestParam{
		// invalid param: board
		{Input: bfsParam{board: [][]int{{1, 1}, {1, 1, 1}, {1, 1, 1}}, endPos: []int{3, 2}}, Expect: -1, ShouldEqual: true},
		{Input: bfsParam{board: [][]int{{1, 1, 1}, {1, -4, 1}, {1, 1, 1}}, endPos: []int{1, 3}}, Expect: -1, ShouldEqual: true},
		// invalid param: endpos
		{Input: bfsParam{board: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, endPos: []int{-1, 1}}, Expect: -1, ShouldEqual: true},
		{Input: bfsParam{board: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, endPos: []int{1, -1}}, Expect: -1, ShouldEqual: true},
		{Input: bfsParam{board: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, endPos: []int{1, 4}}, Expect: -1, ShouldEqual: true},
		{Input: bfsParam{board: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, endPos: []int{4, 2}}, Expect: -1, ShouldEqual: true},
		// unreachable point
		{Input: bfsParam{board: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, -1}}, endPos: []int{2, 2}}, Expect: -1, ShouldEqual: true},
		{Input: bfsParam{board: [][]int{{1, 1, 1}, {1, -1, -1}, {1, -1, 1}}, endPos: []int{2, 2}}, Expect: -1, ShouldEqual: true},
		//
		{Input: bfsParam{board: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, endPos: []int{2, 2}}, Expect: 4, ShouldEqual: true},
		{Input: bfsParam{board: [][]int{{1, 1, 1, -1}, {-1, -1, 1, 1}, {1, 1, 1, -1}, {1, -1, -1, 1}, {1, 1, 1, 1}}, endPos: []int{4, 3}}, Expect: 11, ShouldEqual: true},
	}
	test.Test("测试Bfs", t, params, func(input interface{}) interface{} {
		param := input.(bfsParam)
		return Bfs(param.board, param.endPos)
	})
}
