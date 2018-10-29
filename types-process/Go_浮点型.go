package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明变量并赋值
	var f1 float32 = 3.14
	fmt.Println("f1 == ", f1)

	// 自动推导
	f2 := 1.2456468
	fmt.Println("变量f2的类型是:", reflect.TypeOf(f2))
	fmt.Println("f2 = ", f2)
}
