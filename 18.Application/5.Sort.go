package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{1, 6, 4, 23, 7, 8, 5, 234, 547, 769, 3, 434, 6, 8, 21}
	xs := []string{"Hello", "James", "This", "is", "Hobo", "Scientist", "and", "not", "Dr.Scientist"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)

	fmt.Println(xi)

	sort.Strings(xs)

	fmt.Println(xs)
}
