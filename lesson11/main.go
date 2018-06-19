package main

import "fmt"

func main() {
	// 数组类型: 元素类型，长度
	var a1 [3]int
	//var a2 [4]int
	//a1 = a2 类型不同的，不能赋值
	var a2 [3]int
	a1 = a2
	fmt.Println(a1)
	fmt.Println(a1 == a2)
	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 10: 100} // 可扩容，动态
	fmt.Println(c)

	type stu struct {
		name string
		age  int
	}

	s := stu{
		name: "xxx",
		age:  23,
	}
	d := [3]stu{
		s,
		{"yyy", 25}, // 注意这个逗号！这种方式初始化更方便些，键可以不写
	}
	fmt.Println(d)

	// 最高纬度可以省略，低的不行！
	e := [...][3]int{
		{1, 2},
		{3, 4},
		{5, 6, 7},
	}
	fmt.Println(e)
	fmt.Println(len(d))
	fmt.Println(len(e))
	fmt.Println(cap(e))

	// 数组指针 指针数组
	x, y := 10, 20
	z := [...]*int{&x, &y}
	p := &z
	fmt.Println(*p)
	fmt.Println(&x, &y)

	f := [2]int{11, 22}
	var g [2]int
	g = f // 这里是会复制数据的(全拷贝)！
	fmt.Printf("x: %p, %v \n", &f, f)
	fmt.Printf("x: %p, %v \n", &g, g)
	//test(f) // 这里调用时也会有复制(全拷贝) // 传指针就不会了
	test(&f) // 这里调用时也会有复制(全拷贝) // 传指针就不会了
	fmt.Println(f)
}

func test(x *[2]int) {
	fmt.Printf("x: %p, %v \n", &x, x)
	x[0] = 111
}
