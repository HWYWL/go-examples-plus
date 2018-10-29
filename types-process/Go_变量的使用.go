package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 默认初始化变量为0
	var a int

	fmt.Println("a == ", a)

	// 赋值
	a = 5
	fmt.Println("a == ", a)

	// 初始化并赋值
	var b = 10
	fmt.Println("b == ", b)

	// 自动推导类型
	c := 20
	fmt.Println("变量c的类型是:", reflect.TypeOf(c))

}
