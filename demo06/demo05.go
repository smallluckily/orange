package main

import "fmt"

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scanln(&age)
	switch age {
	case 18:
		fmt.Println("是个刚成年的大小伙子")
	case 19:
		fmt.Println("已经成年了")
	default:
		fmt.Println("xxxxx")

	}

}
