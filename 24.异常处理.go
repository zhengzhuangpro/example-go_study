package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

// 中断程序执行
// Go语言中没有异常机制，但有panic和recover机制
// func init() {
// 	_, err := os.ReadFile("123123")
// 	if err != nil {
// 		fmt.Println("读取文件失败", err)
// 		// 中断程序执行, 只用在初始化函数中
// 		panic(err)
// 	}
// }

// recover 捕获panic抛出的异常
func read() {
	var list []int = []int{1, 2, 3}
	fmt.Println(list[3])
}

// 自定义异常
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

// Go语言中的异常处理机制
func main1() {
	defer func() {
		err := recover() // 捕获异常
		if err != nil {  // 判断是否有异常
			fmt.Println("读取文件失败", err)
			fmt.Printf(string(debug.Stack())) // 打印堆栈信息, 方便调试
		}
	}()

	read()
}

// Go语言中的异常处理机制
func main() {
	// error 是go的一等公民

	a, err := div(10, 1)
	if err != nil {
		// 处理异常
		fmt.Println(err)
		return
	}

	// 正常逻辑
	fmt.Println(a)

	main1()

	fmt.Println("程序正常结束")
}
