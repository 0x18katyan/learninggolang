package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	area := s.length * s.length
	return area
}

func (c circle) area() float64 {
	area := math.Pi * (c.radius * c.radius)
	return area
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("The area of %T is %v.\n", s, s.area())
}

func main() {
	s1 :=
		square{
			length: 4,
		}
	c1 := circle{radius: 4}

	info(s1)
	info(c1)
}
