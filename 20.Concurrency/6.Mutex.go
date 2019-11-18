package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			//runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)

}
