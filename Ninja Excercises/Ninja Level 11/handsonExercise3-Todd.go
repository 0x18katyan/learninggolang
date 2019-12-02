package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "Need more Coffee",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("Foo ran - ", e)

	// code below is not going to work since e is of type error and error has no info
	//	fmt.Println("Foo ran - ", e.info)

	/* Assertion is when you know that a value is of certain type and so you specify it like in the below code
	e.(customErr).info is asserting that e is of type customErr and thus taking .info out of it */
	fmt.Println("Foo ran - assertion", e.(customErr).info)
}
