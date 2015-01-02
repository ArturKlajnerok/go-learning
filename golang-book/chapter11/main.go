package main

import "fmt"
import "github.com/ArturKlajnerok/go-samples/golang-book/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
