package main

import (
	"fmt"
	"unsafe"
)

/*
type node struct { // 自动调整位置，进行对齐补齐，让所占空间最小
	id int32
	next *node
	//_ int
	age int8
	tel int16
}
*/

type user struct {
	name string
	age  int
}

type attr2 struct {
	owner int
	perm  int
}

type file2 struct {
	name string
	attr2 // 匿名属性 组合方式
}

func main() {
	u := user{
		name: "xww",
		age:  11,
	}
	fmt.Println(u)
	u2 := user{"xww2", 22} // 不指定类型的，需要全部写完
	fmt.Println(u2)

	u3 := struct {
		name string
		age  int
	}{
		"123",
		456,
	}
	fmt.Println(u3)

	type file struct {
		name string
		attr struct {
			owner int
			perm  int
		}
	}

	f := file{
		name: "a.txt",
	}
	// 方式1
	f.attr.owner = 1
	f.attr.perm = 1
	// 方式2
	attr := struct {
		owner int
		perm  int
	}{
		1, 1,
	}
	f.attr = attr
	// 空结构
	var a struct{}
	var b [100]struct{}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
	// 组合的方式
	f3 := file2{
		name: "b.txt",
		attr2: attr2{
			1, 2,
		},
	}
	fmt.Println(f3)
}
