package main

import "fmt"

func main() {
	var condition bool = true

	// if语句
	if condition {
		// do something
	} else {
		// do something
	}

	if condition {
		// do something
	} else if condition {
		// do something
	} else {
		// do something
	}

	// switch语句
	switch condition {
	case true:
		// do something
	case false:
		// do something
	default:
		// do something
	}

	// switch语句可以省略条件
	switch {
	case condition:
		// do something
	case condition:
		// do something
	default:
		// do something
	}

	// switch语句可以用于类型判断
	switch i := 2; i {
	case 1:
		// do something
	case 2:
		// do something
	default:
		// do something
	}

	// ==
	fmt.Println(1 == 1)
	// !=
	fmt.Println(1 != 1)

	// >
	fmt.Println(1 > 1)
	// <
	fmt.Println(1 < 1)
	// >=
	fmt.Println(1 >= 1)
	// <=
	fmt.Println(1 <= 1)
	// &&
	fmt.Println(1 == 1 && 1 == 1)
	// ||
	fmt.Println(1 == 1 || 1 == 1)
	// !
	fmt.Println(!true)
	// % 取余
	fmt.Println(10 % 3)
	// & 与
	fmt.Println(1 & 1)
	// | 或
	fmt.Println(1 | 1)
	// ^ 异或
	fmt.Println(1 ^ 1)
	// << 左移
	fmt.Println(1 << 1)
	// >> 右移
	fmt.Println(1 >> 1)
	// &^ 清零
	fmt.Println(1 &^ 1)
	// ^ 异或
	fmt.Println(1 ^ 1)

}
