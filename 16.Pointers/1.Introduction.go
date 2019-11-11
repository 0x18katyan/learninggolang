package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	fmt.Println("After this:")
	b := test(a)
	fmt.Println(b)

	c := &b

	fmt.Printf("%T %v\n", c, c)

	fmt.Println("After Modification: ")

	d := "Changed String"

	fmt.Printf("%T %v\n", d, d)

	e := "Howdie"
	fmt.Printf("%T %v\n", e, e)

	//e = *&d
	fmt.Println(*&e)
	f := fmt.Sprintf("%v %v\n", &e, *&e)
	fmt.Println(f)
}

func test(a int) string {
	re := fmt.Sprintf("%v\n", &a)
	return re
}
