package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type byAge []person

type byName []person

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a byName) Len() int {
	return len(a)
}

func (a byName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byName) Less(i, j int) bool {
	return a[i].first < a[j].first
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Hobo", 24}
	p3 := person{"Monepenny", 27}
	p4 := person{"J Paye Tei", 420}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(byAge(people))

	fmt.Println(people)

	sort.Sort(byName(people))

	fmt.Println(people)
}
