package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	// 设置种子参数，如果改参数相同，则每次生成的随机数也想同
	rand.Seed(time.Now().UnixNano())

	// 随机产生5个一百以内的随机数
	for i := 0; i < 5; i++ {
		fmt.Println("随机数：", rand.Intn(100))
	}
}
