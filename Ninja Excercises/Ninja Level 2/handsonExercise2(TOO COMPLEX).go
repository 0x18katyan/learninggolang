/*
Using the following operators, write expressions and assign their vallues to variables:
==
<=
>=
!=
<
>
*/

package main

import "fmt"

const (
	_ = iota
	a = 1 << (iota * 1)
	b = 1 << (iota * 1)
	c = 1 << (iota * 1)
	d = a
)

func main() {
	fmt.Println(a, b, c)

	checkEquality(a, a)

	/*


		if b >= a {
			fmt.Printf("b is greater than or equal to a where value of a is %v and b is %v\n", a, b)
		}
	*/
}

func checkEquality(firstInt, secondInt int) {
	var toPrint string
	if firstInt == secondInt {
		//fmt.Println("a is equal to b ")
		toPrint = "a is equal to b"
	} else if firstInt != secondInt {
		toPrint = "a is not equal to b"
	}

	fmt.Println(toPrint)
}

func checkLessthanOrEqual(firstInt, secondInt int) {

	if firstInt <= secondInt {
		fmt.Printf("a is less than or equal to b where value of a is %v and b is %v\n", a, b)
	}

}
