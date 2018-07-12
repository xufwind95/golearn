package main

import (
	"fmt"
	"time"
	"sync"
)

func main2() {
	// channel : 同步数据传递、事件通知(close 监听)
	event := make(chan struct{})
	c := make(chan string)
	go func() {
		defer close(event)
		s := <-c
		fmt.Println(s)
	}()
	time.Sleep(2 * time.Second)
	c <- "hi"
	<-event
}

func main1() {
	c := make(chan int, 3)
	defer close(c)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	c <- 3
	c <- 4
	fmt.Println(<-c)
	fmt.Println(<-c)

	// cap len
	a, b := make(chan int, 1), make(chan int, 3)
	a <- 1
	b <- 1
	b <- 2
	fmt.Println("a:", len(a), cap(a)) // 0, 0 默认情况下长度和容量都为0！ 所以初始化时需要写初始容量！
	fmt.Println("b:", len(b), cap(b))

	if value, ok := <-b; ok {
		fmt.Println(value)
	}
}

func main3() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for { // 循环的第一种方式
			if x, ok := <-c; ok {
				fmt.Println(x)
			} else {
				return
			}
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c) // 这里关了才会给子线程发送信号，使之退出！
	<-done
}

func main4() {
	// range
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		// close后的chan，再往里面发送数据出错，不过可以读取值，还有没有取完的，取出，取完了的，读取对应类型的空值
		defer close(done)
		for x := range c { // 循环的第二种方式，这个是检测到close时退出！
			fmt.Println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c) // 这里关了才会给子线程发送信号，使之退出！
	<-done
}

func main5() {
	a := make(chan int, 3)
	a <- 1
	a <- 2
	close(a)
	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println(<-a)
	//a <- 2 // send on closed channel
}

func main6() {
	//var a chan<- int // 单向通道，接收数据
	//var b <-chan int // 单向通道，发送数据
	//a <- 2
	//<-b

	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	go func() {
		defer wg.Done()
		for x := range recv {
			fmt.Println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)
		for i := 0; i < 3; i++ {
			send <- i
		}
	}()
	wg.Wait()

	d := make(chan int, 2)
	var send2 chan<- int = d
	var recv2 <-chan int = d
	//<-send2 // err
	send2 <- 1
	//recv2 <- 2 //invalid operation: recv2 <- 2 (send to receive-only type <-chan int)
	<-recv2
}

func main7() {
	// select
	var wg sync.WaitGroup
	a, b := make(chan int), make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			var name string
			var x int
			var ok bool
			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}
			if !ok {
				return
			}
			fmt.Println(name, "===", x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)
		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()
	wg.Wait()
}

// 定时器
func main() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C: // 没取到东西是会阻塞的
			fmt.Println("he he")
		}
	}
}
