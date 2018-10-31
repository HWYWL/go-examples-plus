package main

import "fmt"

type Person2 struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

type Student2 struct {
	Person2 //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
}

func main() {
	// 定义变量并赋值
	s1 := Student2{Person2: Person2{name: "校花", sex: 18}, id: 1}
	s1.sex = 'M'
	s1.addr = "我家"

	fmt.Println("s1 == ", s1)
}
