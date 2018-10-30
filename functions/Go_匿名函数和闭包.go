package main

import "fmt"

func main() {
	a := 10
	str := "Hello"

	f1 := func() {
		fmt.Println("a == ", a)
		fmt.Println("str == ", str)
	}

	f1()
}
