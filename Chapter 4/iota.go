package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
	d = iota
	e
	f
)

// iota is used to automatically increment something by one

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
