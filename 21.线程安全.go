package main

import (
	"fmt"
	"sync"
)

var sum2 int
var wg1 sync.WaitGroup
var lock1 sync.Mutex

func add2() {
	lock1.Lock()
	for i := 0; i < 100; i++ {
		sum2++
	}
	lock1.Unlock()
	wg1.Done()
}

func sub2() {
	lock1.Lock()
	for i := 0; i < 100; i++ {
		sum2--
	}
	lock1.Unlock()
	wg1.Done()
}

func main() {
	//协程安全
	//互斥锁
	//sync.Mutex

	wg1.Add(2)

	go add2()
	go sub2()

	wg1.Wait()

	fmt.Println(sum2)
}
