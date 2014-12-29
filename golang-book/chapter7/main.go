package main

import (
	"fmt"
)

func addAll(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func multiReturn() (int, int) {
	return 5, 6
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	//multiple values
	x, y := multiReturn()
	fmt.Println(x, y)

	//variadic functions
	xi := []int{5, 6, 7, 8}
	fmt.Println(addAll())
	fmt.Println(addAll(1, 2))
	fmt.Println(addAll(1, 2, 3, 4, 5))
	fmt.Println(addAll(xi...))

	//closure
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 20))

	i := 0
	increment := func() int {
		i++
		return i
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	//recursion
	fmt.Println(factorial(7))

	defer second()
	first()

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
