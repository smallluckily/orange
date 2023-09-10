package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)
	if age < 14 {
		fmt.Println("你是小学生")
	} else if age < 18 {
		fmt.Println("你是未成年")
	} else if age > 23 {
		fmt.Println("你是毕业大学生")
	}
}
