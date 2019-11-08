package main

import (
	"fmt"
)

func main() {
	x := func() int {
		return 420
	}
	y := x()
	fmt.Println(y)
}
