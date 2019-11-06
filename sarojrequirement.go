package main

import "fmt"

type base1 struct {
	type1 int
}

type base2 struct {
	base1
	type2 string
}

type base3 struct {
	base2
	type3 bool
}

func (i base1) sample() {
	fmt.Println("I am base1.")
}

func (i base2) sample() {
	fmt.Println("I am base2.")
}

func (i base3) sample() {
	fmt.Println("I am base3.")
}

type base interface {
	sample()
}

func main() {

	x := base3{
		base2: base2{
			base1: base1{type1: 13},
			type2: "Ashim Mahara"},
		type3: false,
	}
	//fmt.Printf("%T\n", x)

	y := base(x)

	fmt.Printf("%T\n", x.base1)
	x.base1.sample()
	fmt.Println(y)
}
