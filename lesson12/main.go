package main

import (
	"unsafe"
	"fmt"
)

// slice 内部存储方式
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	s := x[2:5]
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s[0])
	//fmt.Println(s[4]) // err 越界了 index out of range
	s1 := make([]int, 3, 5)
	fmt.Println(s1)

	var a []int
	b := []int{}
	fmt.Println(a == nil, b == nil)
	s1 = append(s1, 1)
	//s2 := make()
	s2 := append(s1, 10)
	s2 = append(s2, 10)
	s2 = append(s2, 10)
	fmt.Println(s2, cap(s2)) // 扩容 容量小的时候每次扩都是乘2
	fmt.Println(s1)          // 这个是不变的

	s3 := x[2:4]
	n := copy(s[2:], s3)
	fmt.Println(n)
	fmt.Println(x)
	fmt.Println(s)
}
