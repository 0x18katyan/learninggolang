package main

import (
	"fmt"
)

type person struct {
	first, last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Printf("I am %v %v; and I have a license to kill.\n", s.first, s.last)
}

func (p person) speak() {
	fmt.Printf("I am %v %v; and I don't have a license to kill.\n", p.first, p.last)
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Printf("%v was passed to bar.\n", h)
}

func main() {
	s1 := secretAgent{
		person: person{"James",
			"Bond",
		},
		ltk: true,
	}
	p1 := person{
		first: "Hobo",
		last:  "Scientist",
	}

	bar(s1)
	bar(p1)

	s1.speak()
	p1.speak()
}
