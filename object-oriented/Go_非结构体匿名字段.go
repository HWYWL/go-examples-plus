package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person4 struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

type Student4 struct {
	Person4 //结构体匿名字段
	int    //基础类型的匿名字段
	mystr
}

func main() {
	s := Student4{Person4{"美女", 'm', 18}, 666, "小会"}
	fmt.Println("s == ", s)

	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr)
}
