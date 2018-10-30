package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a == ", a)

	test(a)

	fmt.Println(a)
}

func test(a [5]int)  {
	for i, data := range a {
		a[i] = data + 10
	}

	fmt.Println(a)
}
