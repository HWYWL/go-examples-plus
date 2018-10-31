package main

import "fmt"

type Person5 struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

type Student5 struct {
	*Person5 //指针类型
	id      int
	addr    string
}

func main() {
	s1 := Student5{&Person5{"美女", 'm', 18}, 666, "我家"}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
}
