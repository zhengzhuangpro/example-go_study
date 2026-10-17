package main

import "fmt"

func main() {
	// 值传递
	// 传递的是值的副本
	// 修改副本不会影响原值

	// 引用传递
	// 传递的是值的引用
	// 修改引用会影响原值

	var name string = "hello"
	var a int = 10
	var b *int = &a
	*b = 20
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println("%p\n", &a)

	Set(&name) // &name传递的是name的地址 *name是name的值
}

func Set(name *string) {
	fmt.Println(name)

	var name2 string = "hello"

	*name = name2

	fmt.Printf(*name)
}
