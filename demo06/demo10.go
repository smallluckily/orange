package main

import "fmt"

func main() {
	age := 12
	var pointer *int = &age
	fmt.Println(*pointer)

	v1 := "qzh"
	v2 := &v1
	fmt.Println(*v2)
	v1 = "qzh1"
	fmt.Println(*v2)
}
