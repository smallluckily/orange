package main

import "fmt"

func main() {
	var age [3]int
	age[0] = 22
	age[1] = 33
	age[2] = 44

	var num = [2]int{2, 3}
	var age1 = [3]int{0: 1, 2: 99}
	fmt.Println(age, age1, num)

	for key, item := range age {
		fmt.Println(key, item)
	}
	for i := 0; i < len(age); i++ {
		fmt.Println(i, age[i])

	}
}
