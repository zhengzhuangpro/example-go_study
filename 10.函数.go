package main

import "fmt"

// 无参数无返回值的函数
func sayHello() {
	fmt.Println("Hello World")
}

// 有参数无返回值的函数
func sayHelloTo(name string) {
	fmt.Println("Hello", name)
}

// 有参数有返回值的函数
func add(a, b int) int {
	return a + b
}

func main() {
	// 函数
	// 函数是一段可重复调用的代码块，用于实现特定的功能
	// 函数的定义
	// func 函数名(参数列表) 返回值列表 {
	// 	函数体
	// 	return 返回值
	// }
	// 函数的调用
	// 函数名(参数列表)

	// 无参数无返回值的函数
	sayHello()

	// 有参数无返回值的函数
	sayHelloTo("Tom")

	// 有参数有返回值的函数
	fmt.Println(add(1, 2))
}
