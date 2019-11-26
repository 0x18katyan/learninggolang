package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	go func() {
		fmt.Println(<-c)
		//close(c)
	}()
	c <- 42
	foo()
}

func foo() {

	c := make(chan int, 10) // Using buffered channels
	c <- 42
	fmt.Println(<-c)

}
