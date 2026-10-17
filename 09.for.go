package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 1+2+3+4+5+6+7+8+9+10
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for range
	for i, char := range "hello" {
		fmt.Println(i, char)
	}

	// 死循环
	// for {
	// 	fmt.Println("hello")
	// 	time.Sleep(time.Second) // 休眠1秒
	// }

	// while 模式
	var sun int
	var i int = 1
	for i <= 100 {
		sun += i
		i++
	}
	fmt.Println(sun)

	// do while 模式
	var sum2 int
	var j int = 1
	for {
		sum2 += j
		j++
		if j == 101 {
			break // 结束循环
		}
	}
	fmt.Println(sum2)

	// for range
	for index, value := range []int{1, 2, 3, 4, 5} {
		fmt.Println(index, value)
	}

	for index, value := range [5]int{1, 2, 3, 4, 5} {
		fmt.Println(index, value)
	}

	for index, value := range map[string]int{"a": 1, "b": 2} {
		fmt.Println(index, value)
	}

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 结束循环
		}
		fmt.Println(i)
	}

	// continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 跳过本次循环
		}
		fmt.Println(i)
	}

	// 99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
