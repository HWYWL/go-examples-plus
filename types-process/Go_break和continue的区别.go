package main

import (
	"time"
	"fmt"
)

func main() {
	i := 0

	for {
		i ++
		
		// 演示 1s
		time.Sleep(time.Second)

		if i == 5 {
			//break //跳出循环，如果嵌套多个循环，跳出最近的那个内循环
			continue //跳过本次循环，下一次继续
		}

		fmt.Println("i = ", i)
	}
}
