package main

import "fmt"

func main() {
	str := "abc"

	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%v] == %c\n", i, str[i])
	}

	// 迭代打印每个元素，默认返回两个值：一个元素位置，一个元素本身
	for i, data := range str {
		fmt.Printf("str[%v] == %c\n", i, data)
	}

	// 丢弃第二个值返回值
	for i := range str {
		fmt.Printf("str[%v] == %c\n", i, str[i])
	}

	// 通过匿名变量丢弃第二个返回值
	for i, _ := range str {
		fmt.Printf("str[%v] == %c\n", i, str[i])
	}
}
