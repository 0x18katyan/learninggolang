package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	go addtoChan(c)

	for {
		v, ok := <-c
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}

	fmt.Println("About to Exit main()")
}

func addtoChan(c chan int) {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				c <- j
				wg.Done()
			}
		}()
	}
	wg.Wait()
	fmt.Println("About to exit addtoChan()")
	defer close(c)
}
