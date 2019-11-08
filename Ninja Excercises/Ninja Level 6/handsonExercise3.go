package main

import (
	"fmt"
)

func main() {
	defer bar()
	fmt.Println("This is the beginning.")
	defer fmt.Println("This is after the end.")
	fmt.Println("This is the end.")
}

func bar() {
	fmt.Println("This is barrrrrr.")
}
