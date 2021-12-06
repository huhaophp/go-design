package simple_factory

import (
	"fmt"
)

type Phone interface {
	GetName() string
	SetName(name string)
}

// HuWeiPhone 华为手机
type HuWeiPhone struct {
	Name string
}

func (h *HuWeiPhone) GetName() string {
	return h.Name
}

func (h *HuWeiPhone) SetName(name string) {
	h.Name = name
}

func NewHuWeiPhone() *HuWeiPhone {
	return &HuWeiPhone{Name: "xiaomi"}
}

// XiaoMiPhone 小米手机
type XiaoMiPhone struct {
	Name string
}

func (x *XiaoMiPhone) GetName() string {
	return x.Name
}

func (x *XiaoMiPhone) SetName(name string) {
	x.Name = name
}

func NewXiaoMiPhone() *XiaoMiPhone {
	return &XiaoMiPhone{Name: "xiaomi"}
}

func GetPhone(phone string) (Phone, error) {
	switch phone {
	case "xiaomi":
		return NewXiaoMiPhone(), nil
	case "huawei":
		return NewHuWeiPhone(), nil
	default:
		return nil, fmt.Errorf("wrong phone type passed")
	}
}