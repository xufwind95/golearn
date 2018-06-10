package main

import "fmt"

// 常量命名
const x, y int = 1, 0x23
const s = "hello"
const c = "aa"

func main() {
	//var d int
	const d = 123
	fmt.Println(d)
	const y = 123 // 在此将全部变量覆盖掉
	{
		const x = "abc"
		fmt.Println(x)
	}
	const (
		x1, y1      = 1, 2
		b      byte = byte(x)
		n           = uint8(12)
	)
	// 枚举的指定方式
	const (
		x2 = iota*10 + 1
		y2
		z2 = 100
		d1
		e  = iota
		f
	)
	fmt.Println(x2, y2, z2, d1, e, f) // 1 11 100 100 4 5

	j := 2
	fmt.Println(&j)

	const k = 2
	//fmt.Println(&k) // 常量不变,不能用地址去弄
	b1 := k + 2
	fmt.Println(b1)
}
