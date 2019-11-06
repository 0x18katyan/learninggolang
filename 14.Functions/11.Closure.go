package main

import (
	"fmt"
)

func main() {

	// The following code creates an instance of incrementor(), a, which has it's scope.
	a := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	/* The following codes creates an instance of incrementor() with it's own scope
	   which is different from the variable a.*/

	b := incrementor()
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}

}
