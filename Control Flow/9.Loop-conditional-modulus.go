package main

import (
	"fmt"
)

func main() {
	for x := 0; x <= 20; x++ {
		if x%2 == 0 {
			fmt.Printf("%v is an even number.\n", x)
		} else {
			y := x % 2
			fmt.Printf("%v is an odd number and when divided by 2, %v is the remainder.\n", x, y)
		}
	}
}
