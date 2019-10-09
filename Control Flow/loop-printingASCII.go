/* Loop from 33 to 122 then print the ASCII equivalent of those */

package main

import (
	"fmt"
)

func main() {
	printASCII()
}

func printASCII() {
	i := 33
	for {
		if i > 122 {
			break
		}
		fmt.Printf("%v in decimal is %+q in ASCII\n", i, i)
		i++
	}
}
