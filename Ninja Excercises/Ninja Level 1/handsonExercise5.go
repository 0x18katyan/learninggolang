package main

import "fmt"

type myType int

var x myType
var y int

func main() {
	fmt.Printf("The type of x is %T\n", x)
	fmt.Println("The value of x is", x)
	x = 42
	fmt.Println("After assigning 42 to the variable x, the value of x is", x)
	//fmt.Printf("The type of x is %T\n", x)
	y = int(x)
	fmt.Printf("The type of y is %T\n", y)
	fmt.Println("The value of y is", y)
}
