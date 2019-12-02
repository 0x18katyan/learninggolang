package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("non-existent-file.txt")

	if err != nil {
		panic(err)
	}

	defer fmt.Println("Program Exited")
}
