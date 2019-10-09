package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Printf("%v\t\t%b\t\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%v\t\t%b\t\t%#x\n", y, y, y)

}
