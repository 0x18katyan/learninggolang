package main

import "fmt"

func main() {
	s := "This is a string."
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Converting the string into bytes then viewing the output
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// Print in UTF-8
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	// Iterating through the string with index and value

	for i, v := range s {
		fmt.Println(i, v)
	}
}
