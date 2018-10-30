package main

import "fmt"

// 定义一个结构体
type Stu struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	//顺序初始化，每个成员必须初始化, 别忘了&
	var p1 *Stu = &Stu{1, "校花", 'm', 18, "我家"}
	fmt.Println("p1 = ", p1)

	//指定成员初始化，没有初始化的成员，自动赋值为0
	p2 := &Stu{name: "小萝莉", addr: "也在我家"}
	fmt.Printf("p2 type is %T\n", p2)
	fmt.Println("p2 = ", p2)
}
