package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is an even number.\n", i)
		} else {
			fmt.Printf("%v is an odd number.\n", i)
		}
	}
}
