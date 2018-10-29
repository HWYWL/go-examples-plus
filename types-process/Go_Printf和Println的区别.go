package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30

	// 自动换行
	fmt.Println("a = ", a)

	// 格式化输出，a将占用%d的位置
	fmt.Printf("a = %d\n", a)


	fmt.Println("a = ", a, ",b = ", b, ",c = ", c)
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}
