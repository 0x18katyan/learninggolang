package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-12354)
	if err != nil {
		fmt.Printf("%T\n", err)
		log.Fatalln(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("Error: Square Root of Negative Number: %v", n)
	}
	return 42, nil
}
