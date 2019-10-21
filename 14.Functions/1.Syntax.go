package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	y := convertToFloat(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
}

func convertToFloat(x int) float32 {
	y := float32(x)
	return y
}
