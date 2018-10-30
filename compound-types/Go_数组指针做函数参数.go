package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// 传递地址会直接改变内存的值
	modify(&a)

	fmt.Println(a)
}

func modify(p *[5]int)  {
	for i, data := range *p {
		(*p)[i] = data + 10
	}
}