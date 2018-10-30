package calc

import "fmt"

// 外部调用首字母必须为大写，全显示public，小写是private
func Test()  {
	fmt.Println("我是其他目录的包")
}
