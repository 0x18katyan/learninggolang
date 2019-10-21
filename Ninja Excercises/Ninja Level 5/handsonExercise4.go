package main

import (
	"fmt"
)

type interests map[string]string

func main() {
	x := struct {
		first          string
		last           string
		threeInterests interests
	}{
		first: "hobo",
		last:  "scientist",
		threeInterests: interests{
			"first":  "coding",
			"second": "literature",
			"third":  "music",
		},
	}
	fmt.Println(x)
	fmt.Println(x.first, x.last)
	fmt.Println(x.threeInterests["first"])
	fmt.Println(x.threeInterests["second"])
	fmt.Println(x.threeInterests["third"])
}
