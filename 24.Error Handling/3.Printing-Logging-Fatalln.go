package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("non-existent-file.txt")

	if err != nil {
		f.Close()
		log.Fatalln(err)
	}

	defer fmt.Println("Program Exited")
}
