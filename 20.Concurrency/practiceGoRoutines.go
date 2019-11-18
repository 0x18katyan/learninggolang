package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	x := 0
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go foo(&x)
	}

	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(x)
}

func foo(p *int) {
	*p++
	wg.Done()
	fmt.Println(*p)

}
