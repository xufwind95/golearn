package main

import (
	"errors"
	"fmt"
)

func check(x int) error {
	if x < 0 {
		return errors.New("x < 0")
	} else {
		return nil
	}
}

func main() {
	/*
	x := 6
	if x > 5 {
		fmt.Println(x)
	}else if x > 7 {
		fmt.Println(x)
	}
	*/
	x := 20
	// err in if
	if err := check(x); err == nil {
		x++
		fmt.Println(x)
	} else {
		fmt.Println(err)
	}
	//fmt.Println(err) // zuo yong yu bu gou
	// has no while do while
	for i := 0; i < 4; i++ {
		if i > 3 {
			break
		}
		fmt.Println(i)
	}

	// map channel slice array ==> range
	var data [3]int = [3]int{10, 20, 30}
	for i, s := range data {
		fmt.Println(i, s*10)
		fmt.Println(i, data[i])
	}

	data2 := [3]string{
		"aa",
		"bb",
		"cc",
	}
	var data2Copy [3]*string
	for i, s := range data2 {
		fmt.Println(&i, &s, i, s)
		data2Copy[i] = &s // 引用临时变量的地址，这样很危险
	}
	for _, s := range data2Copy {
		fmt.Println(*s)
	}
	fmt.Println(data2Copy)

	// goto break contiune
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i > 2 {
			goto end // goto 的地方需要对本语句块可见！
		}
	}
end:
	fmt.Println("for cycle end")

	testswitch()
}

func testswitch() {
	x := 1
	a, b, c := 1, 2, 3
	switch x {
	case a, b:
		fmt.Println("x in (a, b)")
		//fallthrough // 继续向下执行(常规的有匹配是不会向下执行的，不过只管一个case！！！)
		//case b:
		//	fmt.Println("x = b")
	case c:
		fmt.Println("x = c")
	default:
		fmt.Println("all not equal")
	}
}
