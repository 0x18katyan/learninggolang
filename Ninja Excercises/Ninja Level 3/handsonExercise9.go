package main

import (
	"fmt"
)

func main() {
	var favSport string
	favSport = "badminton"
	switch favSport {
	case "badminton":
		fmt.Println("His favorite sport is badminton.")
	default:
		fmt.Println("His favorite sport is not badminton.")
	}
}
