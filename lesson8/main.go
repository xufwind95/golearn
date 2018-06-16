package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 100
	p := &a
	fmt.Println(&a)
	fmt.Printf("%p, %v \n", &p, p)
	test(p)
	test2("aaa", 1, 2, 3, 4, 5)

	var a1 []int = []int{1, 2, 3, 4}
	test2("aaa", a1...)
	test2("aaa", a1[:]...)
	fmt.Println(reflect.TypeOf(a1))
	b1 := a1[:]
	fmt.Println(reflect.TypeOf(b1))

	c, _ := test3(1)
	fmt.Println(c)
}

func test(x *int) {
	fmt.Printf("%p, %v\n", &x, x)
}

func test2(s string, a ...int) { //slice
	fmt.Printf("%T, %v\n", a, a)
}

func test3(a int) (age int, name string) {
	b := a / 3
	//age := a / 3 // 这里不能再用这个变量了
	return b, "xww"
}
