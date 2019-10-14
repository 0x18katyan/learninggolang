package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println("The initial length of the slice x is", len(x))
	fmt.Println("The initial capacity of the slice x is", cap(x))
	for i := 0; i < 10; i++ {
		x[i] = i
	}
	fmt.Println(x)
	x = append(x, 100)
	x = append(x, 100)
	fmt.Println(x)
	fmt.Println("The length of the slice x, post-append, is", len(x))
	fmt.Println(`When you cross the capacity of a slice using make then the initial
capacity of the slice is doubled as can be seen below. The initial capacity was 10 and
when I appened two more values to it, the capacity was doubled and became 20.`)
	fmt.Println("The capacity of the slice x, post-append, is", cap(x))
}
