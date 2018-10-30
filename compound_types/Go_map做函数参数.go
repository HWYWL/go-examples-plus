package main

import "fmt"

func main() {
	// map[key]val
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}
	fmt.Println("m = ", m)

	test1(m)
	fmt.Println(m)
}

func test1(m map[int]string)  {
	m[2] = "Hello"
}
