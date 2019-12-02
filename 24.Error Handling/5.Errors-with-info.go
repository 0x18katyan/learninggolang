package main

import (
	"errors"
	"fmt"
	"log"
)

var errMath = errors.New("Error: Square Root of Negative Number")

func main() {
	fmt.Printf("%T\n", errMath)

	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errMath
	}
	return 42, nil
}
