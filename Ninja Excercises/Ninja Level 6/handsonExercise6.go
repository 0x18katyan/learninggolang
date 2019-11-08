package main

import (
	"fmt"
)

func main() {
	x := func() int {
		return 420
	}()
	fmt.Println(x)
}
