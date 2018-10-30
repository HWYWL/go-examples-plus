package main

import "fmt"

func main() {
	MyFunc11(5, 6, 7, 8, 9, 12)
	MyFunc111("校花",18, 5, 6, 7, 8, 9, 12)
}

func MyFunc11(args ...int)  {
	for i, data := range args {
		fmt.Printf("第 %d 参数为 %v\n", i, data)
	}
}

// 前面两个为必填参数
func MyFunc111(name string, age int, args ...int)  {
	for i, data := range args {
		fmt.Printf("%s 今年 %d 岁，第 %d 参数为 %v\n", name, age, i, data)
	}
}
