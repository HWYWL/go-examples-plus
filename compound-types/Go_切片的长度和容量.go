package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	// 从下表0开始取，到下边3结束（不包括下标为3的数据），长度为5（不能超过原有数据长度）
	s := a[0: 3: 5]

	fmt.Println("s == ", s)
	fmt.Println("len(s) = ", len(s)) //长度  3-0
	fmt.Println("cap(s) = ", cap(s)) //容量  5-0

	// 从下表0开始取，到下边4结束（不包括下标为4的数据），长度为5（不能超过原有数据长度）
	l := a[0: 4: 5]

	fmt.Println("l == ", l)
	fmt.Println("len(s) = ", len(l)) //长度  4-0
	fmt.Println("cap(s) = ", cap(l)) //容量  5-0
}
