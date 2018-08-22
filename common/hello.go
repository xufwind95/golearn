package common

import "fmt"

func SayHello() {
	fmt.Println("hello")
}

func init() {
	fmt.Println("init")
}

func init() { // 只有这个函数能重名
	fmt.Println("init2")
}
