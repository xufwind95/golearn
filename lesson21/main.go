package main

import (
	"sync"
	"fmt"
	"time"
)

func main1() {
	// mutex
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(2)
	a := 0
	go func() {
		defer wg.Done()
		for i := 0; i < 5000; i++ {
			m.Lock()
			a += 1
			m.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5000; i++ {
			m.Lock()
			a = a + 1
			m.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(a)
}

type data struct {
	sync.Mutex
	a int
}

func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {
		//d.Lock() 不允许重复加锁！
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	// 上面的方法不用 指针的话，这里是不会达到想要的结果的！因调方法时会复制对象！！！
	var wg sync.WaitGroup
	wg.Add(2)
	var d data
	go func() {
		defer wg.Done()
		d.test("read")
	}()
	go func() {
		defer wg.Done()
		d.test("hello world")
	}()
	wg.Wait()
}
