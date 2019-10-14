package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 113, 69, 420)
	fmt.Println(x)
	//	y := append(x, 123, 83458, 112399, 12413, 23569, 423520)
	//	fmt.Println(y)
	y := []int{10, 20, 30, 40}

	x = append(x, y...)

	fmt.Println(y)
	fmt.Println(x)

}
