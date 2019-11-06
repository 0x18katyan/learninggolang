package main

import (
	"fmt"
)

func main() {
	f := func(x interface{}) {
		fmt.Println(x)
	}
	fmt.Printf("%T\n", f)
	f(32)
}
