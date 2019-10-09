package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		i++
		if i > 100 {
			break
		}
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
	/* Please do not use Eternal Loops without specifying a condition like I did
		 when first writing this program.
	   While using Eternal Loops it is wise to use an outer boundary in conjunction
	   of a break statement since even if the inner boudary is crossed there is
	   no way for a program to break from the loop.  */
}
