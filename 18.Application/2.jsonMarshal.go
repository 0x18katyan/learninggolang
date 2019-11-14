package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "Ashim",
		Last:  "Mahara",
		Age:   24,
	}
	p2 := person{
		First: "Hobo",
		Last:  "Scientist",
		Age:   24,
	}
	people := []person{
		p1,
		p2,
	}

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}
