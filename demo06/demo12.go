package main

import "fmt"

func add(n1 int, n2 int) (int, bool) {
	result := n1 + n2
	return result, true
}

func main() {
	result1, flag := add(3, 7)
	fmt.Println(result1, flag)
}
