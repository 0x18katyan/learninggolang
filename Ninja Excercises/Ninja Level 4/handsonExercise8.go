package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{}
	m["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	m["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}
	for i := range m {
		fmt.Printf("Currently at index %v:\n", i)
		for j := range m[i] {
			fmt.Printf("At index %v of the slice of key %v lies:\t%v\n", j, i, m[i][j])
		}
		fmt.Println("\n")
	}
	othermain()
}

func othermain() {
	m := map[string][]string{}
	m["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	m["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}
	for k, v := range m {
		fmt.Printf("Currently at index %v\n", k)
		for k1, v1 := range v {
			fmt.Printf("At index %v of the slice of key %v lies:\t%v\n", k1, k, v1)
		}
		fmt.Println("\n")
	}

}
