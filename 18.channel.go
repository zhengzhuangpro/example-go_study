package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int)     // 声明并初始化一个长度为0的信道
var nameChan = make(chan string)   // 声明并初始化一个长度为0的信道
var doneChan = make(chan struct{}) // 声明用于关闭的信道

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 结束购物\n", name)

	moneyChan <- money // 向信道中写入数据
	nameChan <- name   // 向信道中写入数据

	wait.Done() // 任务完成
}

// channel
func main() {
	var wait sync.WaitGroup // 等待组
	startTime := time.Now()

	wait.Add(3) // 添加任务

	// 主线程结束，协程也会结束
	go pay("张三", 2, &wait)
	go pay("李四", 3, &wait)
	go pay("王五", 4, &wait)

	go func() {
		defer close(doneChan)  // 关闭信道
		defer close(moneyChan) // 关闭信道
		defer close(nameChan)
		wait.Wait() // 等待组中所有任务完成
	}()

	var moneyList []int
	var nameList []string

	// 从信道中读取数据
	var chanFun = func() {
		for {
			select {
			case money := <-moneyChan: // 从信道中读取数据
				moneyList = append(moneyList, money)
			case name := <-nameChan: // 从信道中读取数据
				nameList = append(nameList, name)
			case <-doneChan: // 从信道中读取数据
				return
			}
		}
	}

	chanFun()
	fmt.Printf("计算结果：%v\n", moneyList)
	fmt.Printf("计算结果：%v\n", nameList)
	fmt.Printf("耗时：%v\n", time.Now().Sub(startTime))
}
