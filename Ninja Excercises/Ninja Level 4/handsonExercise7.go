package main

import (
	"fmt"
)

func main() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Heloooooo, James"}

	x := [][]string{jb, mp}

	for _, v1 := range x {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
