package main

import "fmt"

func main() {
	var ch byte
	var str string

	// 单引号，只能有一个字符
	ch = 'a'
	fmt.Println("ch == ", ch)

	// 双引号，可以有一个或多个字符
	str = "a"
	fmt.Println("str == ", str)

	// 字符串可以通过索引操作
	str = "Hello Go"
	fmt.Printf("str[0] == %c, str[1] == %c\n", str[0], str[1])
}
