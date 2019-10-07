/*
Write a program that prints a number in decimal, binary and hex
*/

package main

import (
	"fmt"
)

func main() {
	var x = 42
	fmt.Printf("%v\t\t%b\t\t%#X\n", x, x, x)
}
