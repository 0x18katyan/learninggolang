package main

import (
	"fmt"
)

func main() {
	hello()
	foo()
}

func hello() {
	fmt.Println("Hello World!")
}

func foo() {

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			response := fmt.Sprintf("The number %d is an even number.", i)
			fmt.Println(response)
		}
	}
}
