package main

import "fmt"

func main() {
	fmt.Println("请输入麻花藤：")
	var name string
	fmt.Scanln(&name)
	if name == "麻花藤" {
		fmt.Println("真聪明")
	} else {
		fmt.Println("你是傻逼么")
	}
}
