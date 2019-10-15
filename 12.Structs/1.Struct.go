package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	p1 := person{
		firstName: "James",
		lastName:  "Bond",
	}

	p2 := person{
		firstName: "Miss",
		lastName:  "Moneypenny",
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
}
