package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)
	for v := range c {
		fmt.Println(v)
	}
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
	/* You need to close the channel once you're done writing to it so that the other functions
	reading from it do not throw an error and know that the channel has been closed.
	*/
}
