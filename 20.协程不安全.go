package main

import (
	"fmt"
	"sync"
)

var sum1 int
var wg sync.WaitGroup

func add1() {
	//协程不安全
	for i := 0; i < 100; i++ {
		sum1++
	}
	wg.Done()
}

func sub1() {
	//协程不安全
	for i := 0; i < 100; i++ {
		sum1--
	}
	wg.Done()
}

func main() {
	//协程不安全
	//协程是轻量级的线程，由操作系统调度，由程序员控制
	//协程是用户态的，线程是内核态的
	//协程是并发的，线程是并行的
	//协程是异步的，线程是同步的
	//协程是抢占式的，线程是协作式的
	wg.Add(2)

	go add1()
	go sub1()

	wg.Wait()

	fmt.Println(sum1)
}
