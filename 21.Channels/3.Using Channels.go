package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int)

	wg.Add(1)
	go send(c)
	wg.Add(1)
	go receive(c)

	fmt.Println("About to Exit")

	/* wg.Add(1) is below here because I have wg.Done() in receive(), would've worked without waitgroup
	if I had used another goroutine for send and a normal statement for recieve because the execution
	will be blocked until all the values in the channel have been dealt with.
	*/

	wg.Add(1)
	go send(c)
	wg.Add(1)
	receive(c)
	wg.Wait()

	// Below is without using waitgroups

	fmt.Println("--------")
	fmt.Println("Without WaitGroup")
	go send2(c)
	receive2(c)

	//	fmt.Println("--------")
	//	fmt.Println("Infinite Send and Receive.")
	//	go infinitesend(c)
	//	infinitereceive(c)
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	wg.Done()
}

func send2(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func receive(c <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	wg.Done()
}

func receive2(c <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func infinitesend(c chan<- int) {
	i := 0
	for {
		c <- i
		i++
	}
}

func infinitereceive(c <-chan int) {
	for {
		fmt.Println(<-c)
	}
}
