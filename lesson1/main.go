package main

import "fmt"

var x = 100

func main() {
	//var x int
	//var y float32
	//var z string
	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(z)

	//var a, s = 1000, "sad" // 自动推断类型
	//fmt.Println(a)
	//fmt.Println(s)
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(s))

	//var (
	//	a int
	//	b string
	//	c float
	//)

	r := 1 //简短命名,这种需要赋值才合法
	fmt.Println(r)

	x := 3 // 声明一个新的变量，和上面那个x没有关系，二者作用域也不一样
	fmt.Println(x)

	x, y := 1, 2
	fmt.Println(x, y)
}
