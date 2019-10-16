package main

import "fmt"

func main() {
	type person struct {
		first, last string
		iceCream    []string
	}
	p1 := person{
		first:    "Ashim",
		last:     "Mahara",
		iceCream: []string{"chocolate", "vanilla", "cream"},
	}
	p2 := person{
		first:    "Hobo",
		last:     "Scientist",
		iceCream: []string{"cold", "juicy", "dark"},
	}
	mapPerson := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	for _, v := range mapPerson {
		fmt.Println("________")
		fmt.Printf("My First Name is %v.\n", v.first)
		fmt.Printf("My Last Name is %v.\n", v.last)
		for _, f := range v.iceCream {
			fmt.Printf("One of my favorite Ice Cream flavors is: %v.\n", f)
		}
	}
}
