package main

import "fmt"

func main() {
	test(10)
}

func test(a int)  {
	if a == 1 {
		fmt.Println("a == ", a)
		return
	}

	// 递归调用
	test(a -1)
	fmt.Println("a == ", a)
}
