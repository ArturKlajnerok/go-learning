package main

import (
	"fmt"
)

func one(xPtr *int) {
	*xPtr = 1
}

func zero(x int) {
	x = 0
}

func zeroPtr(xPtr *int) {
	*xPtr = 0
}

func square(x *float64) {
	*x = *x * *x
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)

	zeroPtr(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

	s := 1.5
	square(&s)
	fmt.Println(s)
}
