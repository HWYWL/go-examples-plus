package main

import "fmt"

func main() {
	// defer延迟调用，在main函数结束前调用
	defer fmt.Println("bbbbbbb")
	defer test1()
	fmt.Println("aaaaaa")

	defer fmt.Println("dddddddddddddddd")
}

func test1()  {
	fmt.Println("xxxxxxxxxxxxxxxxx")
}
