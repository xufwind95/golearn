package main

import (
	"fmt"
	"github.com/xufwind95/golearn/lesson7/pkg"
)

func main() {
	sayHello() // 无需前置声明函数
	sayHello2()
	a := sayHello
	b := sayHello2
	a = b
	a()
	if a == nil {
		fmt.Println("nil")
	}

	var a1 *int = sayHello3()
	fmt.Println(a1, *a1)

	pkg.Simle() // 用大小写来区分可见性
}

func sayHello() {
	fmt.Println("hello world")
	//func aa() {}
	a := func() { fmt.Println("aaa") }
	a()
}

func sayHello2() {
	fmt.Println("hello world")
	a := func() { fmt.Println("aaa") }
	a()
}

func sayHello3() *int {
	a := 100 // 栈  堆
	fmt.Println(&a)
	return &a // 使用逃逸技术直接将之放到堆中去
}
