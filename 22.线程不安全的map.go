package main

import "fmt"

var maps = map[int]string{}

func main() {
	//map是线程不安全的，在并发情况下使用map会报错
	//map的key和value都是引用类型，在并发情况下可能会出现数据错乱的情况
	//可以使用sync.Map来代替map，sync.Map是线程安全的
	//sync.Map的key和value都是interface{}类型，可以存储任意类型的数据

	go func() {
		for {
			maps[1] = "a"
		}
	}()

	go func() {
		for {
			fmt.Println(maps[1])
		}
	}()

	select {}

}
