package main

import (
	"design/simple_factory"
	"fmt"
)

func main()  {
	phone,err := simple_factory.GetPhone("xiaomi")
	if err != nil {
		panic(err)
	}
	fmt.Println(phone.GetName())
}
