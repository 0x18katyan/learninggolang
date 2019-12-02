package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	log.SetOutput(f)

	f2, err := os.Open("non-existent-file.txt")

	if err != nil {
		log.Println("Error:", err)
	}

	defer f2.Close()

	fmt.Println("Check log.txt in this directory for any errors.")
}
