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

// 有多个返回值的函数
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// 可变参数的函数
func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// 命名返回值
func divide2(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// 函数作为参数
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

// 函数作为返回值
func adder() func(int) int {
	sum := 0 // 这里使用了闭包
	return func(x int) int {
		sum += x // 这里使用了闭包
		return sum
	}
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

	// 有多个返回值的函数
	quotient, remainder := divide(10, 3)
	fmt.Println(quotient, remainder)

	// 可变参数的函数
	fmt.Println(sum(1, 2, 3, 4, 5))

	// 命名返回值
	quotient, remainder = divide2(10, 3)
	fmt.Println(quotient, remainder)

	// 函数作为参数
	fmt.Println(apply(add, 1, 2))

	// 函数作为返回值
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// 匿名函数
	var getName = func() string {
		return "Tom"
	}
	fmt.Println(getName())

	// 闭包
	f := adder()
	fmt.Println(f(1), f(2), f(3))

}
