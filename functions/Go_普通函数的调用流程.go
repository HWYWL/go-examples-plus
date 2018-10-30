package main

import "fmt"

func main() {
	funca(3)
}

func funca(a int)  {
	a -= 1
	fmt.Println("a == ", a)
	funcb(a)
}

func funcb(b int)  {
	b -= 1
	fmt.Println("b == ", b)
	funcc(b)
}

func funcc(c int)  {
	c -= 1
	fmt.Println("c == ", c)
}
