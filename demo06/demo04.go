package main

import "fmt"

func main() {
	fmt.Println("请输入两个值用于求和：")
	var a int
	var b int
	fmt.Scanln(&a, &b)
	fmt.Println("a:", a, "b:", b)
	var number int
	number = a + b
	fmt.Println("number:", number)
}
