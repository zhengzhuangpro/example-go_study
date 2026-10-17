package main

import "fmt"

func main() {
	// 1.打印字符串
	fmt.Println("Hello World")

	// 2.打印字符串，不换行
	fmt.Print("Hello World")

	// 3.打印字符串，并换行
	fmt.Println("Hello World")

	// 4.打印字符串，并换行
	fmt.Printf("Hello World\n")

	// 5.打印字符串，并换行
	fmt.Println("Hello World", "Hello Go")

	// 6.%v打印变量
	fmt.Printf("Hello %v\n", "World")

	// 7.%T打印变量类型
	fmt.Printf("%T\n", "World")

	// 8.%d打印整数
	fmt.Printf("%d\n", 10)

	// 9.%f打印浮点数
	fmt.Printf("%f\n", 10.0)

	// 10.%t打印布尔值
	fmt.Printf("%t\n", true)

	// 11.%b打印二进制
	fmt.Printf("%b\n", 10)

	// 12.%o打印八进制
	fmt.Printf("%o\n", 10)

	// 13.%x打印十六进制
	fmt.Printf("%x\n", 10)

	// 14.%X打印十六进制大写
	fmt.Printf("%X\n", 10)

	// 15.%c打印字符
	fmt.Printf("%c\n", 'a')

	// 16.%s打印字符串
	fmt.Printf("%s\n", "Hello World")

	// 17.%p打印指针
	a := 10
	fmt.Printf("%p\n", &a)

	// 18.%e打印科学计数法
	fmt.Printf("%e\n", 10.0)

}
