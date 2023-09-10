package main

import "fmt"

func main() {
	var age = []int{1, 2, 3, 4}
	age[0] = 99
	age[1] = 999
	fmt.Println(age)
	fmt.Println(len(age), cap(age))
	v1 := make([]int, 3, 5)
	v1[0] = 22
	v1[1] = 33
	v1[2] = 34
	fmt.Println(v1)
	fmt.Println(len(v1), cap(v1))
}
