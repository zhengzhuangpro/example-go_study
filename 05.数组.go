package main

import "fmt"

func main() {
	// 声明数组
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// 声明数组并初始化
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	// 声明数组并初始化，指定长度
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// 声明数组并初始化，指定长度和值
	arr3 := [3]int{1: 2, 2: 3}
	fmt.Println(arr3)

	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for index, value := range arr3 {
		fmt.Println(index, value)
	}

	for _, value := range arr3 {
		fmt.Println(value)
	}

	for index := range arr3 {
		fmt.Println(index)
	}

	// 下标
	fmt.Println(arr3[0])

	// 长度
	fmt.Println(len(arr3))

	// 修改元素
	arr3[0] = 10
	fmt.Println(arr3)

	// 多维数组
	arr4 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr4)

	// 多维数组遍历
	for _, arr := range arr4 {
		for _, value := range arr {
			fmt.Println(value)
		}
	}
}
