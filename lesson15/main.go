package main

import (
	"fmt"
	"sync"
)

//import "fmt"
//
//type Stu struct {
//	name string
//	age  int
//}
//
//type N int
//
//func main() {
//	stu := Stu{
//		"123",
//		456,
//	}
//	fmt.Println(stu.setAge(33))
//
//	var a N = 23
//	fmt.Println(a.toString())
//}
//
//// 方法
//func (s *Stu) setAge(a int) int {
//	s.age = a
//	return s.age
//}
//
//func (n N) toString() string {
//	return fmt.Sprintf("%#x", n)
//}

type N int

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("v: %p, %v\n", n, *n)
}

type Stu struct {
	age int
	people
}

func (s Stu) value() {
	s.age = 11
}

func (s *Stu) pointer() {
	s.age = 11
}

type people struct {
	name string
}

func (p *people) getname() string {
	return p.name
}

type data struct {
	sync.Mutex
	buf [1025]byte
}

func main() {
	var a N = 23
	//a.pointer()
	a.value()

	var s1 Stu
	s1.age = 22
	//s1.value() // 22 类型方法:调用方法时,是将原对象复制一份后在用复制品作为第一个默认参数调用的方法
	s1.pointer() // 11 指针方法:不会复制一份对象，传地址
	fmt.Println(s1.age)

	var s2 *Stu
	s2 = &Stu{age: 22, people: people{name: "xxx"}}
	//s2.pointer() //11
	s2.value() // 22 类型方法: 调用方法时,是将原对象复制一份后在用复制品作为第一个默认参数调用的方法
	fmt.Println(s2.age)

	fmt.Println(s2.getname()) // 匿名字段的方法可以直接拿过来

	d := data{}
	d.Lock() // 直接用匿名字段的方法，这里是把锁直接拿过来用的方式
	// do your work
	d.Unlock()
}
