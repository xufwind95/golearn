package main

import (
	"fmt"
	"math"
)

type Student struct {
	Age  int
	Name string
}

func main() {
	/*
	bool
	byte
	int 8 16 32 64 int8 int16
	uint
	float 8 16 32 64
	complex64 128
	uintptr
	string
	array

	struct
	function
	interface

	map 字典
	slice 切片
	channel 管道
	*/
	s := "abc"
	fmt.Println([]byte(s)) // 转换为字节数组
	//b := false
	//stu = &Student{} (返回地址)
	stu := new(Student) //返回一个指针(地址)
	fmt.Println(*stu)

	fmt.Println(math.MinInt8, math.MaxInt16, math.MaxUint16)

	m := make(map[int]string)
	m[1] = "123"
	m[2] = "345"
	fmt.Println(m)

	// 相当于可变数组
	//sl := []int{1, 2, 3, 4, 5}
	//sl = sl[2: 4]
	//fmt.Println(sl)
	sl := make([]int, 0) // 初始不分配容量
	sl = append(sl, 12)
	sl = append(sl, 23)
	sl[1] = 16
	fmt.Println(sl)

	// 消息队列/管道, 2 为容量 如果不设置容量会导致阻塞！
	cl := make(chan int, 2)
	cl <- 1
	fmt.Println(<-cl)

	/*
	make包括了new, 这里只用new是不够的，还有些内部属性需要弄
	p := new(map[int]string)
	m1 := *p
	m1[1] = "123" //assignment to entry in nil map
	fmt.Println(m1)
	*/

	a := 10
	b := byte(a)
	//b := uint8(a)
	fmt.Println(b)

	// 自定义类型, 除操作符以外自定义类型不会继承基础类型的其他东西
	type color int
	var red color
	red = 1 // 赋值不是使用
	var green color
	green = 2
	fmt.Println(red, green)
	/*
	red1 := int(1)
	red1 = red // 不能隐式转化
	fmt.Println(red1)
	*/
}
