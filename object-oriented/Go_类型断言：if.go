package main

import (
	"fmt"
)

type Student10 struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1                    //int
	i[1] = "hello go"           //string
	i[2] = Student10{"美女", 2} //Student

	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Printf("i[%d] 类型为int, 内容为%v\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("i[%d] 类型为string, 内容为%v\n", index, value)
		} else if value, ok := data.(Student10); ok == true {
			fmt.Printf("i[%d] 类型为Student10, 内容为%v\n", index, value)
		}
	}
}
