package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println("After deleting from a slice:")
	x = append(x[:3])
	fmt.Println(x)
	othermain()
}

func othermain() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9, 10}

	// Appending a slice x and slice of a slice y to a slice z.
	z := append(x, y[0:2]...)
	// Always remember to use '...' at the end of a slice while appending.
	fmt.Println(z)

}
