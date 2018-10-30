package main

import "fmt"

func main() {
	var t complex128
	t = 2.5 + 3.14i
	fmt.Println("t == ", t)

	// 自动类型推导
	t2 := 3.5 + 4.6i
	fmt.Println("t2 == ", t2)

	// 通过内建函数，获得实部和虚部
	fmt.Println("real(t2) = ", real(t2), ", imag(t2) = ", imag(t2))
}
