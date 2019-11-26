package main

import "fmt"

func main() {
	c := make(chan int) // Make a channel which holds data of type int

	// c <- 42 // Add 42 to that channel

	go func() {
		c <- 42
		// fmt.Println(<-c) // Won't Work
	}()
	fmt.Println(<-c)
	foo()
	//	bar()
}

func foo() {
	testChannel := make(chan int, 12) // the second argument to the make specifies the size of the buffer
	for i := 0; i < 100; i++ {
		go func() {
			testChannel <- i
		}()
		/* go func() {
			fmt.Println(<-testChannel)
		}() */
		fmt.Println(<-testChannel)
	}
}

func bar() {
	c := make(chan int, 1)
	c <- 1
	c <- 2
	//fmt.Println(<-c)
}
