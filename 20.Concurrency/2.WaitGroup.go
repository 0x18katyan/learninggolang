package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Add(1) // You add a thing to the wait group with this command

	go foo()
	bar()

	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait() // This command helps to wait until everything added to the wait group has been done and removed
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // Then remove it from the wait group with this command when it's done
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
