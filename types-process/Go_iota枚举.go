package main

import "fmt"

func main() {
	// iota常量自动生成器，每一行自动加+
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	// iota遇到const重置为0
	const d int = iota
	fmt.Println("d == ", d)

	// 可以只写一个iota，下面的也会累加赋值
	const (
		a1 = iota //0
		b1
		c1
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

	// 如果是同一行，值都一样
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i = %d, j1 = %d, j2 = %d, j3 = %d, k = %d\n", i, j1, j2, j3, k)
}
