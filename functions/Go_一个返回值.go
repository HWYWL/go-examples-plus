package main

import "fmt"

func main() {
	fmt.Println("MyFunc22 == ", MyFunc22())
	fmt.Println("MyFunc222 == ", MyFunc222())
	fmt.Println("MyFunc2222 == ", MyFunc2222())
}

func MyFunc22() int {
	return 666
}

// 给返回值去一个变量名，go推荐写法
func MyFunc222() (result int) {
	return 666
}

// 给返回值去一个变量名，go推荐写法
func MyFunc2222() (result int) {
	result = 666
	return
}
