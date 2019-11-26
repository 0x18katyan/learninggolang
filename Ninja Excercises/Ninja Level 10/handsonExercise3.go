package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	go gen(c)

	receive(c)

	fmt.Println("About ot exit.")
}

func gen(c chan int) <-chan int {
	//	c := make(chan int)
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
