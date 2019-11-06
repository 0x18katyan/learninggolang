package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("I am being printed from inside an Anonymous Function.")
	}()
	func(x int) {
		fmt.Printf("The number is %v, and this is being printed using an Anonymous Function.\n", x)
	}(32)

}
