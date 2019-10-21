package main

import (
	"fmt"
)

//Defer is used to execute something after the execution of the function ends
func main() {
	defer foo()
	bar()
	fmt.Println("The main() function exited after this, so foo should be printed next.")
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
