package main

import "fmt"

func init() {

	fmt.Println("init1函数执行")
	return
}

func main() {
	// init函数和defer函数
	// init函数：在包级别声明的函数，没有参数和返回值，不能被显式调用，在包初始化时自动执行

	// init函数示例
	fmt.Println("init函数执行")

	// defer函数：在函数执行完毕后，延迟执行指定的函数或代码块，通常用于释放资源、关闭文件等操作
	// defer函数示例
	defer fmt.Println("defer函数执行1")
	defer fmt.Println("defer函数执行2")
	fmt.Println("main函数执行")

	defer func() {
		fmt.Println("defer函数执行3")
	}()

	// defer fmt.Println(a) // defer函数可以访问函数的局部变量，但是不能修改
	// var a int = 10

}
