package main

import "fmt"

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	a, b := 10, 20

	// 两个值交换
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Printf("a = %d, b = %d\n", a, b)

	// 两个值交换
	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

	a1, b1, c1 := test()
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

	// 匿名变量，数据将会被丢弃
	a2, _, c2 := test()
	fmt.Printf("a2 = %d, c2 = %d\n", a2, c2)
}
