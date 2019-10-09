package main

import "fmt"

const (
	/*
		w = 2015 << (iota * 1
		x = w << (iota * 1)
		y = x << (iota * 1)
		z = y << (iota * 1)
	*/
	w = (2015 + iota)
	x
	y
	z
)

func main() {
	fmt.Println(w, x, y, z)
}
