package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)
	receive(even, odd, quit)
	fmt.Println("Quitting function main()")

	foo()
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	//close(even)
	//close(odd)
	//quit <- 0
	close(quit)
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("The value received from the even channel is:", v)
		case v := <-odd:
			fmt.Println("The value received from the odd channel is:", v)
		case v, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok:", v, ok)
				return
			} else {
				fmt.Println("from comma ok:", v)
			}
		}
	}
}

func foo() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}
