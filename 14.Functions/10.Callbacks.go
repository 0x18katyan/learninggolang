package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	sAll := sum(x...)
	sEven := even(sum, x...)
	sOdd := odd(sum, x...)
	fmt.Println("This is the sum of all:", sAll)
	fmt.Println("This is the sum of even numbers:", sEven)
	fmt.Println("This is the sum of odd numbers:", sOdd)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

/* In the code below, even takes 2 arguments, first is a function f which takes
xi which is a variadic parameter of type int and then returns an int.
The other one is vi, which is a variadic parameter of type int.

We then initialize a variable yi of type slice of int then append any values
in vi which is an even number.

At the last step, we unfurm yi, which is a slice of int, into the first argument
and return the output which should be of type int. */

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)

}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
