package main

import "fmt"

//定义一个结构体类型
type Girl struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	s1 := Girl{1, "校花", 'm', 18, "我家"}
	s2 := Girl{1, "校花", 'm', 18, "我家"}
	s3 := Girl{2, "校花", 'm', 18, "我家"}
	fmt.Println("s1 == s2 ", s1 == s2)
	fmt.Println("s1 == s3 ", s1 == s3)

	//同类型的2个结构体变量可以相互赋值
	var tmp Girl
	tmp = s3
	fmt.Println("tmp = ", tmp)
}
