package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)
	//  cs := make(chan<- int) //problematic code
	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)
	fmt.Printf("-----------\n")
	fmt.Printf("cs\t%T\n", cs)

	foo()
}

func foo() {
	cr := make(chan int)
	//	cr := make(<-chan int) // problematic code
	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)
	fmt.Printf("-------\n")
	fmt.Printf("cr\t%T\n", cr)

}
