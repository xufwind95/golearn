package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)
	m["a"] = 3
	fmt.Println(m)
	// 可以用第二个值判定是否有这个键，没有的也不会报错,会用第二个参数来标识
	v, ok := m["c"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no such key")
	}
	m["abc"] = 123
	fmt.Println(len(m))
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	// 不能并发的读写，需要加锁才行！
	m1 := make(map[string]int)
	go func() {
		for {
			m1["a"] = 1
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			_ = m1["b"]
			time.Sleep(time.Microsecond)
		}
	}()

	select {}
}
