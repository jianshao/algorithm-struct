package trie

import (
	"testing"

	"github.com/jianshao/datastruct/utils/test"
)

type phoneIdentifyParam struct {
	operate string
	phone   string
	val     interface{}
}

func Test_PhoneIdentify(t *testing.T) {
	phone := InitPhoneIdentify()
	params := []test.TestParam{
		{Input: phoneIdentifyParam{operate: "add", phone: "18897653478", val: 1234}, Expect: nil, ShouldEqual: true},
		{Input: phoneIdentifyParam{operate: "get", phone: "18897653478"}, Expect: 1234, ShouldEqual: true},
		{Input: phoneIdentifyParam{operate: "remove", phone: "18897653478"}, Expect: nil, ShouldEqual: true},
		{Input: phoneIdentifyParam{operate: "get", phone: "18897653478"}, Expect: nil, ShouldEqual: true},
	}
	test.Test("phone identify", t, params, func(input interface{}) interface{} {
		param := input.(phoneIdentifyParam)
		switch param.operate {
		case "add":
			return phone.AddPhoneNum(param.phone, param.val)
		case "remove":
			return phone.RemovePhoneNum(param.phone)
		case "get":
			val, _ := phone.SearchPhoneNum(param.phone)
			return val
		default:
			return nil
		}
	})
}
