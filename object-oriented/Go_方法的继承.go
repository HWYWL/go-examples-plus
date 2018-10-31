package main

import "fmt"

type Person7 struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//Person类型，实现了一个方法
func (tmp *Person7) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生，继承Person字段，成员和方法都继承了
type Student7 struct {
	Person7 //匿名字段
	id     int
	addr   string
}

func main() {
	var p = Student7{Person7{"校花", 'M', 18}, 1, "我家"}
	p.PrintInfo()
}
