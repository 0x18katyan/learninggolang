package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	y := x[0:4]
	fmt.Println(y)
	fmt.Println(x)

	// Just fun
	for i := 0; i < 4; i++ {
		fmt.Println(x[i])
	}

	/* Doesn't work with negatives, have to figure out another way of getting the
	   last element */

	fmt.Println(x[-1])
}
