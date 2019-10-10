package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Printf("%v is an even number.\n", i)
		case i%2 != 0:
			fmt.Printf("%v is an odd number.\n", i)
		}
	}
}
