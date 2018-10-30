package main

import "fmt"

//定义在函数外部的变量是全局变量
var a int

func main() {
	{
		// 局部变量，离开大括号分配的空间会被释放
		i := 10
		fmt.Println("I == ", i)
	}

	// 此局部变量a不是全局变量a,安装就近原则打印33
	a := 33
	fmt.Println(a)

	test2()
}

func test2() {
	// 只定义未赋值 所以默认值为0
	fmt.Println("test a = ", a)
}

