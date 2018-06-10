package main

import (
	"fmt"
	"strconv"
)

// 变量大写的就可以导出,反之不能,方法函数都是这样区分的
// 大小写决定公有私有!
type Student struct {
	Name string
	age  int
}

func main() {
	fmt.Println("aaa")
	// _ 作为占位符使用
	x, _ = strconv.Atoi("12")

}
