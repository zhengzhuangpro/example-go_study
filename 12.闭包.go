package main

import (
	"fmt"
	"time"
)

// 闭包/高阶函数
func awaitAdd(awaitSecond int) func(...int) int {
	// time.Sleep(time.Duration(awaitSecond) * time.Second)
	return func(args ...int) (sum int) {
		time.Sleep(time.Duration(awaitSecond) * time.Second)
		for _, v := range args {
			sum += v
		}
		return sum
	}
}

func main() {
	// 闭包：函数和其引用环境组合在一起
	// 闭包可以捕获其所在作用域中的变量，即使函数在其原始作用域之外执行，也可以访问这些变量
	t1 := time.Now()
	add := awaitAdd(2)(1, 2, 3, 4, 5)
	t2 := time.Since(t1)
	fmt.Println(add, t2)

}
