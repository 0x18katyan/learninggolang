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

import (
	"fmt"
)

func main() {
	a := (42 == 42)
	b := (123 <= 1214)
	c := (123897 >= 123897)
	d := (12390 != 12390)
	e := (12391 < 129038)
	f := (987098 > 981203)
	fmt.Println(a, b, c, d, e, f)
}
