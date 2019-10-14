package main

import "fmt"

func main() {
	m := map[string]int{
		"ashim":  5,
		"mahara": 6,
	}
	fmt.Println("Before deletion:", m)
	delete(m, "ashim")
	fmt.Println("After deletion:", m)

	// If you delete a non-existent key, there will be no errors at all like
	delete(m, "non-existent")

	m["age"] = 23
	fmt.Println("After adding key age and value 23 to the map:", m)

	/* _,ok is also know as comma ok idiom
	   "_" is used to throw away everything before ok because if you use other identifiers
	  and don't use it, the compiler is gonna complain about it */

	if _, ok := m["age"]; ok {
		delete(m, "age")
		fmt.Println("After deleting key age while using the comma ok idiom, ", m)
	}
}
