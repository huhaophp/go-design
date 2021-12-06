package simple_factory

import "testing"

func TestType1(t *testing.T) {
	api, _ := GetPhone("xiaomi")
	s := api.GetName()
	if s != "xiaomi" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api, _ := GetPhone("huawei")
	s := api.GetName()
	if s != "huawei" {
		t.Fatal("Type1 test fail")
	}
}
