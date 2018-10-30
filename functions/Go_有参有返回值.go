package main

import "fmt"

func main() {
	max, min := MaxMin(15, 20)
	fmt.Printf("最大值为：%v, 最小值为：%v", max, min)
}

func MaxMin(a ,b int) (max, min int) {
	if a > b {
		max = a
		min = b
	}else {
		max = b
		min = a
	}

	return
}
