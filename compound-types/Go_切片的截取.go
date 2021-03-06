package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//[0:len(array):len(array)] 不指定容量和长度一样
	s1 := array[:]
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	s2 := array[3:6:7] //a[3], a[4], a[5]   len = 6-3=3    cap = 7-3=4
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	s3 := array[:6] //从0开始，去6个元素，容量也是6， 常用
	fmt.Println("s3 = ", s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	s4 := array[3:] //从下标为3开始，到结尾
	fmt.Println("s4 = ", s4)
	fmt.Printf("len = %d, cap = %d\n", len(s4), cap(s4))
}
