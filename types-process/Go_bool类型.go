package main

import "fmt"

func main() {
	// 声明变量没有赋值的时候，默认为false
	var a bool
	fmt.Println("a == ", a)

	a = true
	fmt.Println("a == ", a)

	// 声明变量并赋值
	var b bool = true
	fmt.Println("b == ", b)

	// 自动类型推导
	c := true
	fmt.Println("c == ", c)
}
