package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "abc\x61"
	fmt.Println(s)
	// nil ""
	var s2 string
	//fmt.Println(s2 == nil) // err
	fmt.Println(s2 == "")

	s3 := "abc\\ndef" // 转意方式1
	s4 := `abc\ndef`  // 转意方式2
	fmt.Println(s3, s4)

	s4 = "ab" + "cd"
	fmt.Println(s4 == "abcd")
	fmt.Println(s4 > "abc")

	s5 := "abcde"
	fmt.Printf("%c", s5[1]) // 可以进行索引
	//s5[1] = "f" // err
	s6 := s5[1:3]
	fmt.Println(s5, s6)
	s7 := s5[:4]
	fmt.Println(s7)
	fmt.Printf("%#v \n", (*reflect.StringHeader)(unsafe.Pointer(&s5)))
	fmt.Printf("%#v \n", (*reflect.StringHeader)(unsafe.Pointer(&s6)))
	fmt.Printf("%#v \n", (*reflect.StringHeader)(unsafe.Pointer(&s7)))

	for i := 0; i < len(s5); i++ {
		fmt.Printf("%c", s5[i])
	}

	for i, c := range s5 {
		fmt.Printf("%d, %c  ", i, c)
	}
	// 字符串间接修改，不过性能超低！
	b := []byte(s5)
	fmt.Println(b)
	b[0] = 65
	s8 := string(b)
	fmt.Println(s8)

	//s9 := make([]byte, 100)
	//bytes.NewBuffer()
}
