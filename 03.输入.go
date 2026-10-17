package main

import "fmt"

func main() {
	// 1. 从标准输入读取
	fmt.Println("请输入你的名字：")
	var name string
	var age int
	fmt.Scan(&name)
	fmt.Println("name:", name, "age:", age)
}
