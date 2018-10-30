package main

import "fmt"

func main() {
	srcSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6}

	copy(dstSlice, srcSlice)

	// [1 2 6 6 6]
	fmt.Println(dstSlice)
}
