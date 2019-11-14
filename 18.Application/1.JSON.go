package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
    {"Name" : "Hobo Scientist", "Moniker": true},
    {"Name" : "Ashim Mahara", "Moniker": false}
    ]`)

	type Person struct {
		Name    string
		Moniker bool
	}

	var people []Person
	err := json.Unmarshal(jsonBlob, &people)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("%+v\n", people)
}
