package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	length float64
}

func main() {
	square := square{length: 2.7}
	triangle := triangle{base: 7, height: 19}
	printArea(square)
	printArea(triangle)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.length * s.length
}
