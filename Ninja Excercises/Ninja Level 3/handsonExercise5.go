package main

import (
	"fmt"
)

func main() {
	for x := 10; x <= 100; x++ {
		y := x % 4
		fmt.Println(y)
	}
}
