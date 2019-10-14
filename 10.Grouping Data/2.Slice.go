package main

import (
	"fmt"
)

func main() {
	//x := type(values) // composite literal
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)
	othermain()
}

func othermain() {
	x := []string{"I", "am", "batman."}
	fmt.Printf("The type of variable x is %T and the values in variable x are %v.\n", x, x)
	// The above line prints the slice inside of square brackets and I don't know why.
}
