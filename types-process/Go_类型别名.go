package main

import "fmt"

func main() {
	// 给int64起一个别名为bigint
	type bigint int32

	var a bigint
	fmt.Println("a == ", a)

	type (
		long int64
		char byte
	)

	var b long = 100000
	var c char = 'a'

	fmt.Printf("b == %v, c == %c", b, c)
}
