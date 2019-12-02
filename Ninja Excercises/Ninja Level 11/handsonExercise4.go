package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Math Error: Latitude: %v  Longitude: %v  Error: %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-420)

	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		error := sqrtError{
			err:  fmt.Errorf("Number %v is less than 0", f),
			lat:  "10000",
			long: "11000",
		}
		return 0, fmt.Errorf(error.Error())
	}

	return 420, nil
}
