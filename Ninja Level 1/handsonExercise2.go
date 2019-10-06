package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Printf("Variable x is of type %T and has value %v\n", x, x)
	fmt.Printf("Variable y is of type %T and has value %v\n", y, y)
	fmt.Printf("Variable z is of type %T and has value %v\n", z, z)
	fmt.Println(`Because we didn't assign values in the program, all the variables above have Zero Value.`)
}
