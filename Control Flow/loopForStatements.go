package main

import (
	"fmt"
)

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("DONE !")
	eternalLoop()
}

func eternalLoop() {
	fmt.Println("Eternal Loop with Barriers Commences Hence")
	i := 1

	/* If you comment out the IF statement block, the code will run forever which
	   is not really advised since the ctrl+c will not work for some time */

	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}
