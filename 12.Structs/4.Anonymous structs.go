package main

import (
	"fmt"
)

//This works
func main() {
	type person struct {
		first, last string
		age         int
	}
	p1 := person{
		first: "hobo",
		last:  "scientist",
		age:   420,
	}
	fmt.Println(p1)
	othermain()
}

//but

func othermain() {
	fmt.Println("\n\n\nWe are now in othermain().")
	p1 := struct {
		first, last string
		age         int
	}{
		first: "hobo",
		last:  "scientist",
		age:   420,
	}
	fmt.Println(p1)
	fmt.Println("\nThe following code was used for creating an anonymous struct.\n")
	fmt.Println(`p1 := struct {
  first, last string
  age         int
}{
  first: "hobo",
  last:  "scientist",
  age:   420,
}`)
	fmt.Println(`When you are using composite literals to assign values to a previously
undefined struct while creating the struct on the go, the struct is anonymous struct because
it doesn't have a name.`)
}
