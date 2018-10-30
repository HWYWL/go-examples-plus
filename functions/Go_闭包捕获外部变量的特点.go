package main

import "fmt"

func main() {
	a := 10
	str := "Hello"

	func() {
		a = 666
		str = "go"
		fmt.Printf("内部：a = %d, str = %s\n", a, str)
	}()	// 直接调用

	fmt.Printf("内部：a = %d, str = %s\n", a, str)
}
