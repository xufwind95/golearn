package main

import (
	"fmt"
	"time"
	"sync"
)

func task1(s string) {
	fmt.Println("task1 start", s)
	time.Sleep(1 * time.Second)
	fmt.Println("task1 end", s)
}

func task2() {
	fmt.Println("task2 start")
	time.Sleep(2 * time.Second)
	fmt.Println("task2 end")
}

func main() {
	// gorouting = 协程 或者 线程 go runtime
	fmt.Println(time.Now())
	//exit := make(chan int)
	var wg sync.WaitGroup
	/*
		go task1("task1")
		go task2()
	*/
	wg.Add(1)
	go func(s string) { // 用匿名函数启动一个新的线程
		fmt.Println("task3 begin", s)
		time.Sleep(1 * time.Second)
		fmt.Println("task3 end", s)
		//exit <- 1
		wg.Done() //wg.Add(-1) // 这种也可以
	}("task3")
	//time.Sleep(1 * time.Second)
	fmt.Println(time.Now())
	//<-exit
	wg.Wait() // 为0时直接退出，所有一定要把这里搞成0！
}
