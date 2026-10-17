package main

import "fmt"

func main() {
	// 布尔类型
	var a bool = true
	var b bool = false
	fmt.Println(a, b)

	// 整型
	var c int = 1033333
	var d int8 = 127
	var e int16 = 32767
	var f int32 = 2147483647
	var g int64 = 9223372036854775807
	fmt.Println(c, d, e, f, g)

	// 浮点型
	var h float32 = 3.1415926
	var i float64 = 3.1415926
	fmt.Println(h, i)

	// 字符串
	var j string = "hello world"
	fmt.Println(j)

	// uint8
	var k uint8 = 255 // 0~255
	fmt.Println(k)

	// byte 等价于uint8
	var l byte = 'a'
	fmt.Println(l)

	// rune 等价于int32
	var m rune = '中'
	fmt.Printf("%c, %d\n", m, m) // 20013

	// 复数类型
	var n complex64 = 1 + 2i
	fmt.Println(n)

	// 指针类型
	var o *int
	o = new(int)
	fmt.Println(o)

	// 转义字符
	fmt.Println("hello\tworld")
	fmt.Println("hello\nworld")
	fmt.Println("hello\"world\"")
	fmt.Println("hello'world'")
	fmt.Println("hello\\world")
	fmt.Println(`
hello
world
	`)

}
