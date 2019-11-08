package main

import (
	"fmt"
)

func main() {
	x := func() int {
		return 420
	}
	bar(x)
}

func bar(function func() int) {
	x := function()
	fmt.Println(x)
}
