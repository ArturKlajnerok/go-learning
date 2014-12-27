package main

import "fmt"

var (
	a = 5
	b = 10
	c = 15
)

func main() {
	var x1 string
	x1 = "Hello World"
	fmt.Println(x1)

	var x2 string
	x2 = "first "
	fmt.Println(x2)
	x2 = x2 + "second"
	fmt.Println(x2)

	var x3 string = "hello"
	var y3 string = "world"
	fmt.Println(x3 == y3)

	var x4 string = "hello"
	var y4 string = "hello"
	fmt.Println(x4 == y4)

	x5 := 5
	fmt.Println(x5)

	const c string = "Constant"
	fmt.Println(c)

	d := a + b + c
	fmt.Println(d)
}
