package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Ashim":  5,
		"Mahara": 6,
	}
	fmt.Println(m)
	fmt.Println("\n")
	fmt.Println(`Using m["Ashim"] to fetch the corresponding value to the key:`)
	fmt.Println(m["Ashim"])
	fmt.Println("\n")

	fmt.Println(`Using a non-existent key "Non-Existent" as a key in map:`)
	fmt.Println(m["Non-Existent"])
	fmt.Println(`If you use a key which is non-existent in a map, it will return a 0
as can be seen above.`)

	v, ok := m["Non-Existent"]
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println(`If the value of ok in "value, ok := m["Non-Existent"]" is false
then the key doesn't exist in the map like the value above.`)
	// If the value of ok is false then the key doesn't exist in the map.
	checkKey()
}

func checkKey() {
	fmt.Println("I am now in function checkKey()")
	m := map[string]int{
		"Ashim":  5,
		"Mahara": 6,
	}
	fmt.Println(m)

	// Combining the last block in main() with IF statement yields:

	if v, ok := m["Non-Existent-Key"]; ok {
		fmt.Println(`This won't be printed if the key is not in the map. If the key is
indeed present then the value is`, v)
	}
	if v, ok := m["Ashim"]; ok {
		fmt.Println(`This won't be printed if the key is not in the map. If the key is indeed present then the value is`, v)
	}

}
