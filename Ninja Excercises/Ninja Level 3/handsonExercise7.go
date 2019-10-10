package main

import (
	"fmt"
)

func main() {
	for x := 0; x <= 100; x++ {
		y := x % 3
		if y == 0 {
			fmt.Printf("%v is perfectly divisible by 3.\n", x)
		} else {
			fmt.Printf("When %v is divided by 3, %v is the remainder.\n", x, y)
		}
	}
}
