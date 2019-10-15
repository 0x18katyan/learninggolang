package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	fmt.Println("Before deletion:", m)
	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("\n")
		fmt.Printf("We are in the key %v of map m.\n", k)
		for k1, v1 := range v {
			fmt.Printf("We are at index %v with value \"%v\" of key %v.\n", k1, v1, k)
		}
	}
}
