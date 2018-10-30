package main

import "fmt"

func main() {
	var a [10]int

	fmt.Printf("len(a) = %d\n", len(a))

	// 循环赋值
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	fmt.Println(a)

	// 初始化赋值
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println("b == ", b)

	// 推导，没有初始化的元素全为0
	c := [5]int{1, 2, 3}
	fmt.Println("c == ", c)

	// 指定元素赋值
	d := [5]int{2: 5, 4: 10}
	fmt.Println("d == ", d)
}
