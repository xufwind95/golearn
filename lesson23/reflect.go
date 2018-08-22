package main

import (
	"reflect"
	"fmt"
	"unsafe"
)

type X int
type Y int

func main1() {
	/*
	var a X = 100
	t := reflect.ValueOf(a)
	fmt.Println(t.Kind())
	*/
	// TypeOf 获取类型 ValueOf 获取值
	var a, b X = 100, 200
	var c Y = 300
	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)
	fmt.Println(ta == tb, tb == tc, "===", ta, reflect.ValueOf(a))
	fmt.Println(tb.Kind() == tc.Kind()) // Kind 返回go语言的内置类型！
}

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main2() {
	var m manager
	t := reflect.TypeOf(m)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		if f.Anonymous {
			fmt.Println(f.Type, "is a anonymous field:")
			for x := 0; x < f.Type.NumField(); x++ {
				f2 := f.Type.Field(x)
				fmt.Println(f2.Name, f2.Offset, f2.Type)
			}
			fmt.Println("end====================")
		}
	}
}

type A struct {
}
type B struct {
	A
}

func (A) Av()  {}
func (*A) Ap() {}
func (B) Bv()  {}
func (*B) Bp() {}

type User struct {
}
type Admin struct {
	User
}

func (*User) ToString() {}
func (Admin) test()     {}

func main() {
	// 导出的方法(首字母大写)才能被反射找到！
	methods := func(t reflect.Type) {
		fmt.Println(t.Name(), ":", t.NumMethod())
		for i, n := 0, t.NumMethod(); i < n; i++ {
			fmt.Println(t.Method(i).Name)
		}
	}
	var b B
	var u Admin
	fmt.Println("======value interface=======")
	methods(reflect.TypeOf(b))
	methods(reflect.TypeOf(u))
	fmt.Println("======pointer interface=======")
	methods(reflect.TypeOf(&b))
	methods(reflect.TypeOf(&u))
}

type user2 struct {
	name int `field:"name" type:"varchar(50)"` // 注意: 键值对要挨着！！！
	age  int `field:"age" type:"int"`
}

func main4() {
	/*
	var s http.Server
	t := reflect.TypeOf(s) // 可以获得私有变量的数据
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}
	*/
	var u user2
	t2 := reflect.ValueOf(u).Type()
	//fmt.Println(t2)
	//fmt.Println(reflect.TypeOf(u))
	t2 = reflect.TypeOf(u)
	for i := 0; i < t2.NumField(); i++ {
		f := t2.Field(i)
		// 主要用于orm映射
		fmt.Printf("%s: %s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}
}

type user3 struct {
	Name string
	code int
}

func main_() {
	a := 100
	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()
	fmt.Println(va.CanAddr(), va.CanSet())
	fmt.Println(vp.CanAddr(), vp.CanSet())
	fmt.Println(va.Type(), vp.Type())

	p := new(user3)
	v := reflect.ValueOf(p).Elem()
	name := v.FieldByName("Name")
	code := v.FieldByName("code")
	fmt.Println(name.CanAddr(), name.CanSet())
	fmt.Println(code.CanAddr(), code.CanSet()) // 小写(非导出)的不能设置值
	name.SetString("tom")
	fmt.Println(p.Name)

	//code.SetInt(1)
	//fmt.Println(p.code) //

	*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 11
	fmt.Println(p.code)
}
