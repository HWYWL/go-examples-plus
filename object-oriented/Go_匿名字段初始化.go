package main

import "fmt"

type Person1 struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

type Student1 struct {
	Person1 //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
}

func main() {
	// 定义变量并赋值
	var s1  = Student1{Person1{"校花", 'M', 18}, 1, "我家"}
	fmt.Println("s1 == ", s1)

	// 自动推导类型
	s2 := Student1{Person1{"美女", 'M', 18}, 1, "我家"}
	fmt.Println("s2 == ", s2)
}
