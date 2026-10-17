package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 结束购物\n", name)
	wait.Done() // 任务完成
}

// 协程
func main() {
	var wait sync.WaitGroup // 等待组
	startTime := time.Now()

	wait.Add(3) // 添加任务

	// 主线程结束，协程也会结束
	go shopping("张三", &wait)
	go shopping("李四", &wait)
	go shopping("王五", &wait)

	wait.Wait() // 等待组中所有任务完成

	// for {
	// 	if wait == 0 {
	// 		break
	// 	}
	// }

	// time.Sleep(1 * time.Second)

	fmt.Printf("主函数执行结束，耗时：%v\n", time.Since(startTime))
}
