package main

import "fmt"

func main() {
	var name string
	fmt.Println("请输入你的名字")
	fmt.Scanf("%s", &name)
	if name == "qzh" {
		fmt.Println("输入的名字为正确")
	} else {
		fmt.Println("输入的名字为错误")
	}

}
