package main

import (
	"example-go_study/version"
	"fmt"
)

func hello() {
	fmt.Println("hello")
}

var age int = 18

const pi = 3.14

func main() {
	// 1.变量定义
	// 1.1.显式定义
	var a int
	var b float32
	var c string
	var d bool

	c = "hello world"

	f := 1.2 // 自动推导类型

	fmt.Println(a, b, c, d, f)

	hello()

	age := 20

	fmt.Println(age)

	name, gender := "张三", "男"

	age, gender = 18, "女"

	fmt.Println(name, gender)

	var (
		name1   string = "张三"
		gender1 string = "男"
	)

	fmt.Println(name1, gender1)

	// 2.常量定义

	// pi = 3.1415926 // 常量不可修改

	fmt.Println(pi)

	// 外部变量
	fmt.Println(version.Version)
	// fmt.Println(version.commit) // 首字母小写，外部包不可访问

}
