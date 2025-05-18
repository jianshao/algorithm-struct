package trie

import (
	"strconv"
	"testing"

	"github.com/jianshao/datastruct/utils"
	"github.com/jianshao/datastruct/utils/test"
)

type trieParam struct {
	operate string
	val     interface{}
	path    []interface{}
}

const (
	OPERATE_GET    = "get"
	OPERATE_ADD    = "add"
	OPERATE_DELETE = "delete"
)

func Test_Trie(t *testing.T) {
	// 以0-9的数字为路径,如12345
	root := Init(10, func(val interface{}) int {
		v := val.(string)
		if len(v) != 0 {
			return -1
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return -1
		}
		return i
	})
	params := []test.TestParam{
		{Input: trieParam{operate: OPERATE_ADD, path: utils.Arr2InterfaceArr("1234"), val: 1}, Expect: 1, ShouldEqual: true},
		{Input: trieParam{operate: OPERATE_GET, path: utils.Arr2InterfaceArr("1234")}, Expect: 1, ShouldEqual: true},
		{Input: trieParam{operate: OPERATE_DELETE, path: utils.Arr2InterfaceArr("1234")}, Expect: 1, ShouldEqual: true},
		{Input: trieParam{operate: OPERATE_GET, path: utils.Arr2InterfaceArr("1234")}, Expect: nil, ShouldEqual: true},
	}
	test.Test("trie", t, params, func(input interface{}) interface{} {
		param := input.(trieParam)
		if param.operate == OPERATE_ADD {
			if root.Add(param.path, param.val) == nil {
				return 1
			} else {
				return 0
			}
		} else if param.operate == OPERATE_DELETE {
			root.Remove(param.path)
			return 1
		} else if param.operate == OPERATE_GET {
			result, _ := root.Get(param.path)
			return result
		} else {
			return nil
		}
	})
}
