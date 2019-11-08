package main

import (
	"fmt"
)

func main() {
	x := factorialRevamp(5)
	fmt.Println(x)
	y := factorial(5)
	fmt.Println(y)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialRevamp(n int) int {
	factorial := 1
	//	for n := n; n > 0; n-- {
	/* If the value is already initialized then the first section of a for loop
	   which is initialization can be skipped over using ';' */
	for ; n > 0; n-- {
		factorial *= n
	}
	return factorial
}
