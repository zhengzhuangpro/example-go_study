package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Name string `json:"name"`
}

type Person struct {
	Class
	Age int `json:"-"` // json tag - 表示该字段不序列化也就是不转换成json
}

func (s *Person) SayHello() {
	fmt.Println("Hello, my name is", s.Name)
}

func (s *Person) Info() {
	fmt.Println("My name is", s.Name, "and I am", s.Class.Name, "years old", s.Name)
}

func (s *Person) SetAge(age int) {
	s.Age = age
}

func main() {
	// 结构体
	// 结构体是一种聚合数据类型，它是由一系列具有相同类型或不同类型的数据构成的数据集合
	// 结构体是一种用户自定义的数据类型，它可以包含多个不同类型的数据成员
	// 结构体是一种复合数据类型，它可以包含多个不同类型的数据成员
	// 结构体是一种自定义的数据类型，它可以包含多个不同类型的数据成员

	// 定义结构体
	var p Person
	p.Name = "张三"
	p.Age = 18
	fmt.Println(p)

	p.SayHello()

	s1 := Person{
		Age: 20,
		Class: Class{
			Name: "高三",
		},
	}
	s1.SayHello()

	// 结构体指针
	// 结构体指针是一种指向结构体的指针类型，它可以指向一个结构体变量
	// 结构体指针是一种复合数据类型，它可以包含多个不同类型的数据成员
	// 结构体指针是一种自定义的数据类型，它可以包含多个不同类型的数据成员
	// 结构体指针是一种用户自定义的数据类型，它可以包含多个不同类型的数据成员
	s2 := &Person{
		Age: 20,
		Class: Class{
			Name: "高三33",
		},
	}

	s2.SayHello()

	fmt.Println(s2)
	byteData, _ := json.Marshal(s2)
	fmt.Println(string(byteData))
}
