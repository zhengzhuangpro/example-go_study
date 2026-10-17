package main

import "fmt"

func main() {
	// 切片是引用类型
	// 切片底层是一个数组
	// 切片是一个动态数组
	// 切片可以自动扩容
	// 切片不能直接比较，可以使用reflect.DeepEqual()函数比较
	// 切片不能直接修改底层数组，需要通过索引修改
	// 切片可以遍历
	// 切片可以截取
	// 切片可以追加
	// 切片可以删除
	// 切片可以复制
	// 切片可以合并
	// 切片可以反转

	// 切片定义
	var s1 []int
	var s2 = []int{}
	var s3 = make([]int, 0)
	var s4 = []int{1, 2, 3, 4, 5}
	var s5 = s4[1:3]
	var s6 = s4[:3]
	var s7 = s4[3:]
	var s8 = s4[:]

	fmt.Println(s1, s2, s3, s4, s5, s6, s7, s8)

	// 切片长度和容量
	fmt.Println(len(s4), cap(s4))

	// 切片遍历
	for i := 0; i < len(s4); i++ {
		fmt.Println(s4[i])
	}

	// 新增切片元素
	s4 = append(s4, 6)
	fmt.Println(s4)

	// 删除切片元素
	s4 = append(s4[:2], s4[3:]...)
	fmt.Println(s4)

}
