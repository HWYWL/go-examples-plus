package main

import "fmt"

func main() {
	// 多个变量定义，自动推导
	var (
		a = 10
		b = 20.5
	)

	// 多个常量定义
	const (
		c = 10
		d = 10.5
	)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
}
