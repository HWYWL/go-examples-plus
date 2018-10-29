package main

import (
	"fmt"
	"reflect"
)

func main() {
	//变量：程序运行期间，可以改变的量， 变量声明需要var
	//常量：程序运行期间，不可以改变的量，常量声明需要const

	// 两种声明方式
	const a int = 10
	const b = 11.5
	fmt.Println("a = ", a)

	fmt.Println("变量c的类型是:", reflect.TypeOf(b))
	fmt.Println("b = ", b)
}
