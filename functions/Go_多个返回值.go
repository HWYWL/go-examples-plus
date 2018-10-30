package main

import "fmt"

func main() {
	a, b, c := MyFunc3()
	fmt.Printf("a = %d, b = %d, c = %d\n",a ,b ,c)
	a, b, c = MyFunc33()
	fmt.Printf("a = %d, b = %d, c = %d\n",a ,b ,c)
	a, b, c = MyFunc333()
	fmt.Printf("a = %d, b = %d, c = %d\n",a ,b ,c)
}

// 多个返回值
func MyFunc3() (int, int, int) {
	return 1, 2, 3
}

// go推荐写法
func MyFunc33() (a int, b int, c int) {
	a, b, c = 11, 22, 33
	return
}

// go推荐写法
func MyFunc333() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}
