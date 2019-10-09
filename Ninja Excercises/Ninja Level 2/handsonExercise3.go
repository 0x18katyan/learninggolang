package main

import (
	"fmt"
)

//const x = "Thus Spoke Zarathustra"
//const y int = 10

///*
//The above can also be declared using Parens such as
const (
	x     = "Ths Spoke Zarathustra"
	y int = 10
)

//*/
func main() {
	fmt.Println("The value of constant x is", x)
	fmt.Println("The value of typed constant y is", y)
}
