package main

import "fmt"

type Person8 struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//有个学生，继承Person字段，成员和方法都继承了
type Student8 struct {
	Person8 //匿名字段
	id     int
	addr   string
}

//Person类型，实现了一个方法
func (tmp *Person8) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//Student也实现了一个方法，这个方法和Person方法同名，这种方法叫重写
func (tmp *Student8) PrintInfo() {
	fmt.Println("Student: tmp = ", tmp)
}

func main() {
	p := Student8{Person8{"校花", 'M', 18}, 1, "我家"}
	// 调用Student8的方法，先找本作用域，找不到再找继承
	p.PrintInfo()
}
