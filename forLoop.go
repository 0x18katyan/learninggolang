package main

import "fmt"

func main() {
	someFunction()
}

func someFunction() {
	for i := 0; i <= 10; i++ {
		fmt.Println(fmt.Sprintf("Current Value of i is %d.", i))
	}
}
