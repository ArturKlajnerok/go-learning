package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

type Circle struct {
	x float64
	y float64
	r float64
}

var c1 Circle

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return (l + w) * 2
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleAreaLocal(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {

	cp := new(Circle)
	fmt.Println(cp)

	c2 := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c2)

	c3 := Circle{0, 0, 5}
	fmt.Println(c3.x, c3.y, c3.r)
	c3.x = 10
	c3.y = 5

	fmt.Println(circleArea(c3))

	fmt.Println(circleAreaLocal(&c3))

	fmt.Println(c3.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	a := new(Android)
	a.Name = "Loli"
	a.Person.Talk()
	a.Talk()

	fmt.Println(totalArea(&c3, &r))

	m := MultiShape{[]Shape{&c2, &c3, &r}}
	fmt.Println("area:", m.area())
	fmt.Println("perimeter:", m.perimeter())
}
