package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("Variable x is of type %T and value %v, Variable y is of type %T and value %v, and Variable z is of type %T and value %v.", x, x, y, y, z, z)
	fmt.Println(s)

	// Below is the code from the previous handsonExercise2.
	/*	fmt.Printf("Variable x is of type %T and has value %v\n", x, x)
		fmt.Printf("Variable y is of type %T and has value %v\n", y, y)
		fmt.Printf("Variable z is of type %T and has value %v\n", z, z)
	*/

	fmt.Println("\n")
	fmt.Println(`Since we have already declared the variables x,y,z in the other go program named
handsonExercise2.go, the compiler is giving me a warning about it, it's safe to ignore that warning
since I could have used the previous file for this exercise but not so otherwise; kudos to the go-plus package
in Atom.`)
	//	fmt.Println(`Because we didn't assign values in the program, all the variables above have Zero Value.`)
}
