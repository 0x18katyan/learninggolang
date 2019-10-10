package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		for x := 0; x < 3; x++ {
			//fmt.Printf("%U\t\t%c\n", i, i)
			// %c is for converting the decimal to character related to the decimal in unicode.
			// OR using a single 'verb', i.e. %#U
			fmt.Printf("%#U\n", i)
		}
	}

}
