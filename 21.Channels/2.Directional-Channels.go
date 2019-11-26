package main

import (
	"fmt"
)

func main() {
	c := make(chan<- int) // Send only channel // Can only send into the channel and not recieve

	for i := 0; i < 10; i++ {

		go func() {
			c <- i
		}()
		//	fmt.Println(<-c) // Doesn't run
	}

	fmt.Printf("%T\n", c)
	foo()
	//bar()
	test()
	conversion()
}

func foo() {
	c := make(<-chan int, 2) // This is recieve only channel, we cannot write to it
	fmt.Printf("%T\n", c)
}

func bar() {
	c := make(chan int)
	go func() { c <- 1 }()
	fmt.Println(<-c)
}

func test() {
	c := make(chan int, 2)
	c <- 1
	fmt.Println(<-c)
}

func conversion() {
	c := make(chan int)
	cr := make(<-chan int) // recieve only channel
	cs := make(chan<- int) // send only channel

	fmt.Println("-----------")

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// Lets try converion between channels
	fmt.Println("-----------")

	//	fmt.Printf("c\t%T\n", (chan int)(cr)) // Can't convert from specific to general

	// Let's try changing general channel to recieve only
	fmt.Printf("c\t%T\n", (<-chan int)(c)) // Can convert from general to specific

}
