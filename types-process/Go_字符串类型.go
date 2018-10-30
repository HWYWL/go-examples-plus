package main

import "fmt"

func main() {
	// 声明变量并赋值
	var str1 string = "abc"
	fmt.Println("str1 == ", str1)

	// 自动推导
	str2 := "def"
	fmt.Println("str2 == ", str2)

	// 内建函数，len()可以获得字符串的长度
	fmt.Println("len(str2) == ", len(str2))
}
