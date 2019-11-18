package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Number of CPus", runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		x := i
		go func() {
			fmt.Println("This is GoRoutine Number: ", x)
			fmt.Println("Number of GORoutines:", runtime.NumGoroutine())
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Number of GORoutines at the END:", runtime.NumGoroutine())
}
