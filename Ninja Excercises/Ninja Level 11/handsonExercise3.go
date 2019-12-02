package main

import (
	"errors"
	"fmt"
)

type customErr struct {
	err error
}

func main() {
	err := customErr{err: errors.New("This is a custom error")}
	fmt.Printf("%T\n", err)
	fmt.Printf("%T\n", err.err)
	a := foo(err.err)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)
}

func foo(err error) error {
	return fmt.Errorf("Error: %v", err)
}
