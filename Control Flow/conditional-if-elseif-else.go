package main

import (
	"fmt"
)

func main() {
	x := 13
	for x < 20 {
		if x == 13 {
			fmt.Println("The value of variable x is 13.")
		} else if x > 13 {
			fmt.Println("The value of variable x is greater than 13.")
		} else {
			fmt.Println("The value of variable x is less than 13.")
		}
		x++
	}
	dupmain()
}

func dupmain() {
	x := 0
	for x < 20 {
		if x == 13 {
			fmt.Println("The value of variable x is 13.")
		} else if x > 13 {
			fmt.Println("The value of variable x is greater than 13.")
		} else {
			fmt.Printf("The value of variable x is %v.\n", x)
		}
		x++
	}
}
