package main

import "fmt"

func main() {
	fmt.Println("欢迎致电中国移动，您要访问的服务是：1.话费想关;2.业务办理;3.人工服务")
	var number int
	fmt.Scanln(&number)
	if number == 1 {
		fmt.Println("您选择可话费想关服务：1.充值话费;2.查询话费")
		fmt.Scanln(&number)
		if number == 1 {
			fmt.Println("你选择的服务是充值话费话费")
			fmt.Println("请输入您有要充值的话费数额：")
			var money int
			fmt.Scanln(&money)
			fmt.Println("充值成功，您充入了话费：", money)

		} else {
			fmt.Println("你选择的服务是查询话费")
			fmt.Println("请在您的手机上输入10086用于授权查询话费")
			var number1 int
			fmt.Scanln(&number1)
			fmt.Println("查询成功，您的话费是xxxxx")
		}
	} else if number == 2 {

	} else if number == 3 {

	}
}
