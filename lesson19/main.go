package main

import (
	"math"
	"fmt"
	"sync"
)

func count() {
	x := 0
	for i := 0; i < math.MaxInt32; i++ {
		x += 1
	}
	fmt.Println(x)
}

func test1(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			count()
		}()
	}
	wg.Wait()
}

func main() {
	// GOMAXPROCS
	//runtime.GOMAXPROCS(runtime.NumCPU()) // 指定同一时间最多占用多少个处理器(这里是本台机器的CPU的数量)
	/*
	n := runtime.GOMAXPROCS(0) // 有几个CPU用多少个
	fmt.Println(time.Now())
	test2(n)
	fmt.Println(time.Now())
	*/

}
