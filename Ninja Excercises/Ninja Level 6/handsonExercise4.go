package main

import (
	"fmt"
)

type person struct {
	first, last string
	age         int
}

func (person person) speak() {
	fmt.Printf("Hi! My Name is %v %v and I am %v years old.\n", person.first, person.last, person.age)
}

func main() {
	p1 := person{
		first: "Ashim",
		last:  "Mahara",
		age:   24,
	}
	p1.speak()
}
