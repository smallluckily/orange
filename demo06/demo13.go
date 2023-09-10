package main

import "fmt"

type car interface {
	run()
}

type baoshijie struct {
	name string
}
type falali struct {
	name string
}

func (b baoshijie) run() {
	fmt.Printf("%s 速度是70 \n", b.name)
}

func (f falali) run() {
	fmt.Printf("%s 速度是700 \n", f.name)
}

func drive(c car) {
	c.run()
}

func main() {
	var b1 = baoshijie{
		name: "保时捷",
	}
	var f1 = falali{
		name: "法拉利",
	}

	drive(b1)
	drive(f1)

}
