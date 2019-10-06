package main

import (
	"fmt"
)

func main() {
	var x = 13
	fmt.Printf("I am printing the actual value with this : %v \n",x)
	fmt.Printf("I am now printing the go-syntax representation of the value with this : %#v \n",x)
	fmt.Println(`Which didn't make sense to me but I guess I will find it out later`)
	fmt.Printf("I am now going to print the type of x: %T \n", x)
}