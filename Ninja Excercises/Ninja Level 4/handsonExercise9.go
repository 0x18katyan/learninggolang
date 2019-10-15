package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice Cream", "Sunsets"},
	}

	m["scientist_hobo"] = []string{"whatever", "buff jhol momo", "books"}

	for k, v := range m {
		fmt.Println("\n")
		fmt.Printf("We are in %v key of the map.\n", k)
		for k1, v1 := range v {
			fmt.Printf("At index %v of the slice lies \"%v\".\n", k1, v1)
		}
	}

}
