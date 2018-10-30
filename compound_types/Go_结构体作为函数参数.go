package main

import "fmt"

//定义一个结构体类型
type Girl1 struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	s := Girl1{1, "美女", 'm', 18, "我家"}
	fmt.Println("s == ", s)
	test11(s)
	fmt.Println("s == ", s)
	test22(&s)
	fmt.Println("s == ", s)
}

func test11(g Girl1)  {
	g.name = "校花"
	fmt.Println("g == ", g)
}

func test22(p *Girl1)  {
	p.name = "小萝莉"
	fmt.Println("p == ", p)
}
