package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	c := make(chan int)

	go send(ctx, c)
	//	receive(ctx, c)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(c chan int) {
		defer cancel()
		defer wg.Done()
		for v := range c {
			if v == 100 {
				//				cancel() instead of this
				//wg.Done()
				break
			} else {
				fmt.Println(<-c)
			}
		}
	}(c)

	wg.Wait()
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
}

func send(ctx context.Context, c chan<- int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			c <- i
		}
	}
}

func receive(ctx context.Context, c <-chan int) {
	_, cancel := context.WithCancel(ctx)
	for v := range c {
		if v == 100 {
			cancel()
			return
		} else {
			fmt.Println(<-c)
		}
	}
}
