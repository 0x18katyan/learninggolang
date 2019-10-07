package main

import (
	"fmt"
)

func main() {
	var x = 4
	fmt.Printf("%d\t\t%b\n", x, x)

	/* What the following statement does is shift the one's place to another place in the front
	   like 100 becomes 1000 but in binary form, and thus 4 becomes 8 because 100 is 4 and 1000 is 8 in decimal form. */
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)
	//	othermain()
	iotafun()
}

/* func othermain() {
	// Using manual method to shift
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
*/

const (

	/* Basically what the following code does is set the first iota as non-consequential
			then increment the subsequent iota with a measure of 10 in binary. So, kb is initialized as
		  10000000000 then mb is incremented by 1 measure of iota where 1 measure is the kb which is kb which is 10000000000
		  and then adds another 10 zeros after it and same with the gb.
	    Still needs some clearning up.
	*/

	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func iotafun() {

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
