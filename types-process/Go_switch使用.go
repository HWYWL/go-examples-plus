package main

import "fmt"

func main() {
	var num int
	fmt.Printf("请按下电梯楼层：")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("按下的是1楼")
		// 默认包含break
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	case 4:
		fmt.Println("按下的是4楼")
	default:
		fmt.Println("指挥中心，我已离开地球")
	}

	score := 85
	//可以没有条件,case后面可以放条件
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("一般")
	default:
		fmt.Println("其它")
	}
}
