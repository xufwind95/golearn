package main

import (
	"fmt"
	"errors"
)

func test(f func(s string) (int, error)) int {
	a, _ := f("abc")
	return a
}

func test2(s string) (int, error) {
	fmt.Println(s)
	return 2, errors.New("error")
}

func main() {
	/*
	func(s string) {
		fmt.Println(s)
	}("abc")

	result := test(test2)
	fmt.Println(result)

	f, err := os.Open("filename")
	defer f.Close() // 延迟调用
	if err != nil {
		fmt.Println(err)
	}
	var c []byte
	b := bytes.NewBuffer(c)
	content, _ := f.Read(b.Bytes())
	fmt.Println(content)
	*/

	x, y := 1, 2
	defer func(a int) {
		fmt.Println(a, x, y)
	}(11)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			// free resource ...
		}
	}()

	fmt.Println(x, y+1)
	fmt.Println(testdefer())

	// panic recover (try catch)
	panic("database is not connected")
	fmt.Println("test painc")
}

func testdefer() (int) {
	a := 100
	defer func() { // 延迟到最后的一条原子语句之前执行
		fmt.Println("defer")
		a += 100
	}()
	return a
	//temp := a
	//执行 defer
	//return temp
}
