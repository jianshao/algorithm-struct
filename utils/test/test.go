package test

import (
	"reflect"
	"testing"
)

// 需要描述被测试的函数的输入输出
// 对于输出可以直接判断是否相等
// 如何将被测试的函数传入，以及如何描述输入参数？？？
// 将
type TargetFunc func(input interface{}) interface{}

type TestParam struct {
	Input       interface{}
	Expect      interface{}
	ShouldEqual bool
}

const (
	ShouldEqual    = true
	ShouldNotEqual = false
)

func Test(funcName string, t *testing.T, params []TestParam, target TargetFunc) {
	for _, param := range params {
		result := target(param.Input)
		if reflect.ValueOf(result).Equal(reflect.ValueOf(param.Expect)) != param.ShouldEqual {
			t.Errorf("Error occurred when running '%s', Expect: %v Got: %v ShouldEqual: %v", funcName, param.Expect, result, param.ShouldEqual)
		}
	}
}
