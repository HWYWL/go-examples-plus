package main

import "fmt"

//函数也是一种数据类型， 通过type给一个函数类型起名
//FuncType1它是一个函数类型
type FuncType1 func(int, int) int

func main() {
	result := Calc(10, 20, Add1)

	fmt.Println("result == ", result)
}

func Add1(a, b int) (result int) {
	result = a + b
	return 
}

func Calc(a, b int, fTest FuncType1) (result int) {
	result = fTest(a, b)
	return
}