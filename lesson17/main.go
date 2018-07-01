package main

import "fmt"

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct{}

func (d *data) test() {

}

func (d *data) string() string {
	return "123"
}

func main2() {
	//var d data
	//var t tester = &d
	//t.test()
	//fmt.Println(t.string())

	var d data
	var t tester = &d
	pp(t)

	var s stringer = t
	pp(s)

	// var t2 tester = s // 不能子集到父集赋值
}

func pp(a stringer) {
	fmt.Println(a.string())
}

func main() {
	/*
	t := 3
	var f float32 = float32(t)
	fmt.Println(f)

	var t int
	t = 4
	fmt.Println(t)

	var s string
	s = "aa"
	*/

	var i interface{} // 可以承载任何的数据类型!
	i = 4
	fmt.Println(i)
	i = "abc"
	fmt.Println(i)

	m := make(map[string]interface{}, 0)
	m["key1"] = 1
	m["key2"] = "value2"
	fmt.Println(m)

	s := stu{
		"aaa",
		23,
	}
	m["key3"] = s
	fmt.Println(m)

	d := m["key1"].(int) // 类型断言 对接口进行类型断言
	var d1 int
	d1 = d // err
	fmt.Println(d1)

	var s1 stu
	s1, _ = m["key3"].(stu)
	fmt.Println(s1)
	//var s2 *stu
	//s2, _ = m["key3"].(*stu)
	//fmt.Println(s2) // nil 转换是失败的，不过接收第二个值后不会报错(不接收的要报错)

	d2, ok := m["key4"].(int)
	if ok {
		fmt.Println(d2)
	}
}

type stu struct {
	name string
	age  int
}
