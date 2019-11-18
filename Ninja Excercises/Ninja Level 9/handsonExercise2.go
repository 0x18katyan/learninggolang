package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println("My name is", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Ashim",
		last:  "Mahara",
	}

	saySomething(&p1)
	// /	saySomething(p1)
	p1.speak()
}
