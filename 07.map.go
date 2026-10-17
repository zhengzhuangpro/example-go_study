package main

import "fmt"

func main() {
	var userMap map[int]string = map[int]string{
		1: "john",
		2: "peter",
		3: "jane",
	}

	fmt.Println(userMap)
	fmt.Println(userMap[3])
	fmt.Printf("%#v\n", userMap[4]) // 应该返回空字符串

	val, ok := userMap[4]
	fmt.Println(val, ok)

	userMap[4] = "jim" // 添加元素
	fmt.Println(userMap)

	delete(userMap, 4) // 删除key为4的元素
	fmt.Println(userMap)

	var aMap = make(map[string]string)
	aMap["hello"] = "world" // nil map不能赋值
	fmt.Printf("%#v\n", aMap)

}
