package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("non-existent-file.txt")

	if err != nil {
		fmt.Println("Error:", err)
		log.Fatalln("Error:", err)
	}
}
