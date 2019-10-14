package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[3])

	for i, v := range x {
		fmt.Printf("The index is %v and the value at index %v is %v.\n", i, i, v)
	}
}
