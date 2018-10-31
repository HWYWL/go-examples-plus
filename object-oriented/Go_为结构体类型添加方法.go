package main

import "fmt"

type Person6 struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

// 带有接受者的函数叫做方法，传递引用（拷贝一份）
func (tem Person6) PrintInfo()  {
	fmt.Println("tem == ", tem)
}

// 通过一个方法给成员赋值，值引用（直接修改）
func (tem *Person6) SetInfo(name string, sex byte, age int)  {
	tem.sex = sex
	tem.name = name
	tem.age = age
}

type long int64
type char byte

// 只要参数不同即可人认为是不同方法，和方法名无关
func (temp long) test()  {

}

func (temp char) test()  {

}

func main() {
	p := Person6{"美女", 'M', 18}
	p.PrintInfo()

	var p2 Person6
	(&p2).SetInfo("校花", 'M', 18)
	p2.PrintInfo()
}
