package main

import "fmt"

// 接口
type SingInterface interface {
	Sing()
	GetLabe() string
}

type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println(c.Name + "在唱歌")
}

func (c Chicken) GetLabe() string {
	return c.Name
}

type Cat struct {
	Name string
}

func (c Cat) Sing() {
	fmt.Println(c.Name + "在唱歌")
}

func (c Cat) GetLabe() string {
	return c.Name
}

// 定义一个接口，包含一个Sing方法
func sing(c SingInterface) { // 接口作为参数
	// ch, ok := c.(Chicken) // 类型断言
	// fmt.Println(ch, ok)

	// 类型断言
	switch server := c.(type) {
	case Chicken:
		fmt.Println(server)
	case Cat:
		fmt.Println(server)
	default:
		fmt.Println("未知类型")
	}

	c.Sing()
	fmt.Println(c.GetLabe())
}

// 空接口
type EmptyInterface interface{}

func PrintEmptyInterface(v EmptyInterface) {
	fmt.Println(v)
}

func main() {
	fmt.Println("接口")
	// 实现接口
	ch := Chicken{Name: "小鸡"}
	ca := Cat{Name: "小猫"}
	sing(ch) //小鸡在唱歌
	sing(ca) //小猫在唱歌

	PrintEmptyInterface(ch)
	PrintEmptyInterface(ca)
}
