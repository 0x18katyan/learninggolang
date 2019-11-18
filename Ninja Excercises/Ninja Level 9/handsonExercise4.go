package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	counter := 0
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		fmt.Println("The loop variable i is:", i)
		fmt.Println("---------------")
		fmt.Println("Number of GoRoutines right now is: ", runtime.NumGoroutine())
		mu.Lock()
		go func() {
			fmt.Println("Counter during this GoRoutine is: ", counter)
			counter++
			fmt.Println("Counter after increment during this GoRoutine is:", counter)
			wg.Done()
			mu.Unlock()
		}()
		fmt.Printf("Number of GoRoutines at the end of the loop %v is: %v\n", i, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter at the End is: ", counter)
}
