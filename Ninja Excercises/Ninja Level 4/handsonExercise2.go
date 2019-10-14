package main

import (
	"fmt"
)

func main() {
	// I am having confusion between declaration of slice and array
	// In arrays you have to specify the size of array with arr := [2]int but not in slices
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", x)
}
