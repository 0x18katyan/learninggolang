package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Printf("The type of variable x is %T\n", x)
	fmt.Println("The value of variable x is", x)
	boolTest()
}

func boolTest() {
	a := 10
	b := 11
	// "==" operator is used for comparision, returns true if true and false is false
	fmt.Println(a == b)

	c := 13
	d := 13
	fmt.Println(`We can use "==" for comparison. The following code is doing the comparision:
    fmt.Println(c == d)`)
	fmt.Println(c == d)
	fmt.Println(`We can also use "<" or ">" for comparison between values. `)
	fmt.Println(`I am now using the code: "fmt.Println(c > d)" for comparing c and d with values 13 and 13 respectively.`)
	fmt.Println(c > d)
	fmt.Println(`I am now using the code: "fmt.Println(a < b)" for comparing a and b with values 10 and 11 respectively.`)
	fmt.Println(a < b)
	fmt.Println(`You can also use ">=" or "<=" for comparison.
I am now using "fmt.Println(c >= b)" to compare c and d with values 13 and 13 respectively.`)
	fmt.Println(c >= b)
}
