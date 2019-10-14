package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(x)
	y := []string{"Ashim", "Mahara", "Chocolate", "Whiskey"}
	fmt.Println(y)

	// Multi-Dimensional Slice
	xp := [][]string{x, y}
	fmt.Println(xp)

}
