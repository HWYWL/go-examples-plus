package main

import "fmt"

func main() {
	s := "刘强东"

	if s == "王思聪" {
		fmt.Println("左拥右抱")
	} else if s == "刘强东" {
		fmt.Println("奶茶妹妹")
	}else {
		fmt.Println("左手右手")
	}

	// if支持一个初始语句，初始化语句和判断条件以分号分隔
	if a := 10; a == 10 {
		fmt.Println("a == ", a)
	}
}
