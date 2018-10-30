package main

import "fmt"

// 定义一个结构体
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	var s1  = Student{1, "校花", 'M', 18, "我家"}
	fmt.Println("s1 == ", s1)

	// 知道初始化成员，没有初始化的都为0
	s2 := Student{name: "mike", addr: "bj"}
	fmt.Println("s2 = ", s2)
}
