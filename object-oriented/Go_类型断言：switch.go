package main

import "fmt"

type Student11 struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1                    //int
	i[1] = "hello go"           //string
	i[2] = Student11{"校花", 666} //Student

	//类型查询，类型断言
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("i[%d] 类型为int, 内容为%d\n", index, value)
		case string:
			fmt.Printf("i[%d] 类型为string, 内容为%s\n", index, value)
		case Student11:
			fmt.Printf("i[%d] 类型为Student, 内容为name = %s, id = %d\n", index, value.name, value.id)
		}

	}
}
