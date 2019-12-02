package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovering in f", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v is the value from the call panic()", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
