package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//	go func(c chan int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	//  }(c)

	for {
		v, ok := <-c
		if !ok {
			fmt.Println("Breaking out of for loop")
			break
		} else {
			fmt.Println(v)
		}
	}
	fmt.Println("About to Exit main()")
	fmt.Println("--------------------")
	toddsWay()
}

func toddsWay() {
	fmt.Println("Entering Todd's Way")
	a := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			a <- i
		}
		close(a)
	}()
	for x := range a {
		fmt.Println(x)
	}
	fmt.Println("About to Exit todd's way")
}
