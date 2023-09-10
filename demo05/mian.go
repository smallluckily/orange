package main

import "fmt"

// import "fmt"
//
// var number int = 999
//
//	func main() {
//		number := 666
//		fmt.Println(number)
//	}
var number int = 999

func main() {
	number := 666
	if true {
		number := 333
		fmt.Println(number)
	}
	fmt.Println(number)
}
