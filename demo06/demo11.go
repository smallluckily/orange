package main

import "fmt"

func main() {
	type person struct {
		name  string
		age   int
		phone int
	}
	p1 := person{"qzh", 22, 33}
	fmt.Println(p1.name, p1.age, p1.phone)
}
