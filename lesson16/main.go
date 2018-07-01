package main

import "fmt"

type tester interface {
	test()
	string() string
	//aa()
}

type data struct {
}

func (d *data) test() {
	fmt.Println("hello world")
}

func (d *data) string() string {
	return "test interface"
}

func main() {
	//var d data // 类型不一致的不能算实现！ 方法中是 指针类型，这里不是！
	// 不过参数是值类型的话，定义指针是能够接住值类型的方法的！
	d := &data{}
	var i tester
	i = d // d 实现了tester的所有方法
	fmt.Println(i.string())
	i.test()
}
