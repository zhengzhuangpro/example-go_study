package main

import "fmt"

func login() {
	//登录逻辑
	fmt.Println("登录")
}

func register() {
	//注册逻辑
	fmt.Println("注册")
}

func main() {
	// 输入
	fmt.Printf("请选择操作：1.登录 2.注册")
	// 判断
	var choice int
	fmt.Scan(&choice)
	// 执行
	// switch choice {
	// case 1:
	// 	login()
	// case 2:
	// 	register()
	// }

	var funMap = map[int]func(){
		1: login,
		2: register,
	}
	// 执行
	if fun, ok := funMap[choice]; ok {
		fun()
	}

}
