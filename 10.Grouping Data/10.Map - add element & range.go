package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"ashim":  5,
		"mahara": 6,
	}

	fmt.Println(m)
	// You can create a new key and assign it value using the syntax mapIdentified[key] = value
	m["age"] = 23
	fmt.Println(m)

	// The code below iterates over the map and prints the index and value
	for i, v := range m {
		fmt.Println(i, v)
	}
}
