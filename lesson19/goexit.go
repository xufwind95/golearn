package main

import (
	"fmt"
	"runtime"
)

func main() {
	//Goexit 只影响当前线程，且线程中的defer能执行
	exit := make(chan int)
	go func() {
		defer close(exit)
		defer fmt.Println("aaa")
		func() {
			defer func() {
				fmt.Println("bbb", recover() == nil)
			}()
			func() {
				fmt.Println("ccc")
				runtime.Goexit()
				fmt.Println("ccc exit")
			}()
			fmt.Println("bbb exit")
		}()
		fmt.Println("aaa exit")
	}()
	<-exit
	fmt.Println("main exit")
}
