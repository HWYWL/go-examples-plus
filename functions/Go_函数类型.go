package main

import "fmt"

func main() {
	// 声明一个函数类型的变量，变量名叫fTest,并把Add函数赋值
	var fTest FuncType = Add
	result := fTest(10, 20)

	fmt.Println("result == ", result)
}

func Add(a, b int) (result int) {
	result = a + b
	return 
}

//函数也是一种数据类型， 通过type给一个函数类型起名
//FuncType它是一个函数类型
type FuncType func(int, int) int 