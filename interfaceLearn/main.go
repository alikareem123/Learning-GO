package main

import (
	"fmt"
	"math"
)

type helper interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

type Rectangle struct {
	Length float64
	Breadth float64
}

func (r Rectangle) Area() float64{
	return r.Length * r.Breadth
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

func printer(s helper) {
	fmt.Printf("AreaL: %.2f\n ", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	c:= Circle{Radius: 10}
	printer(c)
	r := Rectangle{Length: 10, Breadth: 18}
	printer(r)
}