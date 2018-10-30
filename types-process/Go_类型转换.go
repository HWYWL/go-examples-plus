package main

import "fmt"

func main() {
	// 不能转换的被称为不兼容类型
	flag := true
	fmt.Printf("flag == %t\n", flag)

	// bool 类型不能转换为int类型, 0是false，非0是true

	var ch byte
	ch = 'a'
	t := int(ch)
	fmt.Println("t == ", t)
}
