package main

import (
	"runtime"
	"fmt"
)

func main() {
	//fmt.Println("aaa")
	runtime.GOMAXPROCS(1)
	exit := make(chan int)

	go func() {
		defer close(exit) // 这样不会有内存泄漏 也在函数中可以用 exit <- 1

		go func() {
			fmt.Println("b")
		}()
		fmt.Println("a1")
		runtime.Gosched() // 系统10 ms自动切换一次 一般不要用这个东西
		fmt.Println("a2")
	}()
	<-exit
}
