package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

func printArea(shape Shape) {
	fmt.Printf("Area of %s is: %f\n", shape.Name(), shape.Area())
}

func printPerimeter(shape Shape) {
	fmt.Printf("Perimeter of %s is: %f\n", shape.Name(), shape.Perimeter())
}

type square struct {
	name   string
	length float64
}

type triangle struct {
	name string
	a, b, c float64
}

func (sq square) Area() float64 {
	return sq.length * sq.length
}

func (sq square) Perimeter() float64 {
	return 4.0 * sq.length
}

func (sq square) Name() string {
	return sq.name
}

func (tri triangle) Area() float64 {
	s := (tri.a + tri.b + tri.c) / 2
	area := math.Sqrt(s * (s - tri.a) * (s - tri.b) * (s - tri.c))
	return area
}

func (tri triangle) Perimeter() float64 {
	return tri.a + tri.b + tri.c
}

func (tri triangle) Name() string {
	return tri.name
}

func main() {
	sq := square{
		name:   "square",
		length: 5.0,
	}

	tri := triangle{
		name: "triangle",
		a:    3.0,
		b:    5.0,
		c:    6.0,
	}

	printArea(sq)
	printPerimeter(sq)

	printArea(tri)
	printPerimeter(tri)
}
