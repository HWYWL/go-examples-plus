package main

import fmt "fmt"

func main() {
	var ch byte = 97

	//格式化输出，%c以字符方式打印，%d以整型方式打印
	fmt.Printf("%c, %d\n", ch, ch)

	ch = 'b'
	//格式化输出，%c以字符方式打印，%d以整型方式打印
	fmt.Printf("%c, %d\n", ch, ch)
}
