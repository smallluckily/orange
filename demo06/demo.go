package main

import "fmt"

func main() {
	fmt.Print("请输入你的名字：")
	var name string
	fmt.Scanln(&name)
	if name == "qzh" {
		fmt.Println("您是尊敬的vip客户")
		goto VIP
	} else {
		fmt.Println("您是普通用户")
		goto PUTONG
	}
VIP:
	fmt.Println("这是vip客户通道直接进入通道")
PUTONG:
	fmt.Println("需要排队进入")
}
