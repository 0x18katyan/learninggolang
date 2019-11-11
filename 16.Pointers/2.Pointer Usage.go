package main

import (
	"fmt"
)

func main() {

	a := 0
	fmt.Println("The value of a right now is:", a)
	foo(&a, 50)
	fmt.Println("Now I used the function foo() to modify the value of a using pointers, which now is:", a)
	//	fmt.Println(a)
	bar()
}
func foo(addr *int, toMod int) {
	fmt.Println("The value of addr is:", addr)
	fmt.Println("The address of addr is:", &addr)

	*addr = toMod
}

func bar() {
	x := 420
	y := &x
	fmt.Println(y, *y)
}
