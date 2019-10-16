package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool // license to kill
}

func main() {
	sa1 := secretAgent{
		person: person{"James", "Bond", 69},
		ltk:    true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	//but can also be accessed by using sa1.person.*
	fmt.Println(sa1.person.first, sa1.person.last, sa1.person.age)
	othermain()
}

func othermain() {
	fmt.Println("\n\n")
	fmt.Println("We are now in othermain().")
	type person struct {
		first string
		last  string
		age   int
	}

	type secretAgent struct {
		person
		ltk bool // license to kill
		// i can also use another type of the same identifier here like
		age string
		// but to access this i will have to use absolute traversal (my term)
	}
	sa1 := secretAgent{
		person: person{"James", "Bond", 69},
		ltk:    true,
		age:    "Sixty Nine",
	}

	fmt.Println("This is sa1.age:", sa1.age, "\t\t", "This is sa1.person.age:", sa1.person.age)
}
