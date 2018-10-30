package main

import "fmt"

func main() {
	a := 12
	b := 20

	defer func(a, b int) {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}(30, 25)

	a ,b = 111, 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)
}
