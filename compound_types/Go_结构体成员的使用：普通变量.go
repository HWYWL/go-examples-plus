package main

import "fmt"

// 定义一个结构体
type Student1 struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	var s Student1

	s.id = 1
	s.name = "美女"
	s.sex = 'm' //字符
	s.age = 18
	s.addr = "还在我家"
	fmt.Println("s = ", s)
}
