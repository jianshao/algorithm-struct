package trie

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/jianshao/datastruct/utils"
)

type PhoneIdentify struct {
	root    *Trie
	initted bool
}

func transString2Int(val interface{}) int {
	if reflect.TypeOf(val).Kind() != reflect.String {
		return -1
	}
	// 必须是0-9的数字
	v, err := strconv.Atoi(val.(string))
	if err != nil || v < 0 || v > 9 {
		return -1
	}

	return v
}

// 场景：使用手机号为key对用户进行识别，或存储一些数据
func InitPhoneIdentify() *PhoneIdentify {
	return &PhoneIdentify{
		root:    Init(11, transString2Int),
		initted: true,
	}
}

func (p *PhoneIdentify) AddPhoneNum(phone string, val interface{}) error {
	if !p.initted {
		return errors.New("initialize first please")
	}
	return p.root.Add(utils.Arr2InterfaceArr(phone), val)
}

func (p *PhoneIdentify) SearchPhoneNum(phone string) (interface{}, error) {
	if !p.initted {
		return nil, errors.New("initialize first please")
	}
	return p.root.Get(utils.Arr2InterfaceArr(phone))
}

func (p *PhoneIdentify) RemovePhoneNum(phone string) error {
	if !p.initted {
		return errors.New("initialize first please")
	}
	p.root.Remove(utils.Arr2InterfaceArr(phone))
	return nil
}
