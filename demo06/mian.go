package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("请输入用户的姓名和年龄")
	count, err := fmt.Scan(&name, &age)

	if err == nil {
		fmt.Println(name, age)
	} else {
		fmt.Println(count, err)
	}

}
