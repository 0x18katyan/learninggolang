package main

import (
	"fmt"
)

func main() {
	conditionals()
	initializationScope()
}

func conditionals() {
	if true {
		fmt.Println("This will be printed.")
	}
	if false {
		fmt.Println("This won't be printed.")
		/* The reason that the above code didn't run is go code will only run on true.
		   And, true and false are predeclared constants in Go. */
	}
	if !true {
		fmt.Println("This won't be printed.")
		// The above didn't ran because !true is equal to false.
	}
	if !false {
		fmt.Println("This will be printed.")
		// The above ran because !false is true.
	}
	if 2 == 2 {
		fmt.Println("The statement 2 is equal to 2 is true so this will be printed.")
	}
	if 2 != 2 {
		fmt.Println("The statement 2 is not equal to 2 is false so this will not run.")
	}

	if !(2 == 2) {
		fmt.Println(`This will not be printed since NOT(2 == 2 /*(which is true))*/ is false.`)
	}

	if !(2 != 2) {
		fmt.Println(`This will be printed since !( 2 != 2) is true because 2 != 2 is false and !false is true.`)
	}
}

func initializationScope() {
	if x := 0; x < 100 {
		fmt.Println(x)
	}
}
