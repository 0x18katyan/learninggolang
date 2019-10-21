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

/*
  func (r reciver) identifier(parameters) (return(s)) { code }
 func speak() is only available to the vairables of type secretAgent
 and any variable using the function is assigned variable s for the duration
 of that execution and within the scope of the speak() function.
*/

func (s secretAgent) speak() {
	if s.ltk == true {
		fmt.Printf("I am %v %v, and I have a license to kill.\n", s.first, s.last)
	} else {
		fmt.Println("I am ", s.first, s.last)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{"James", "Bond"},
		ltk:    true,
	}
	sa2 := secretAgent{
		person: person{"Hobo", "Scientist"},
		ltk:    false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()
}
