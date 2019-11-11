package main

import (
	"fmt"
)

type person struct {
	first, last string
	age         int
}

func main() {
	p1 := person{
		first: "Hobo",
		last:  "Scientist",
		age:   24,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(person *person) {
	(*person).first = "Joker"
	fmt.Println(person.first)
	fmt.Println("Here comes the Joker.")
	(*person).first = "Bruce"
	(*person).last = "Wayne is BATMAN!!!"
	fmt.Println(person)
}
