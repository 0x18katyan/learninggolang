package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var counter int64
	counter = 0
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		fmt.Println("The loop variable i is:", i)
		fmt.Println("---------------")
		fmt.Println("Number of GoRoutines right now is: ", runtime.NumGoroutine())
		go func() {
			fmt.Println("Counter during this GoRoutine is: ", atomic.LoadInt64(&counter))
			// counter++
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter after increment during this GoRoutine is:", atomic.LoadInt64(&counter))
			wg.Done()
			runtime.Gosched()
		}()
		fmt.Printf("Number of GoRoutines at the end of the loop %v is: %v\n", i, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter at the End is: ", counter)
}
