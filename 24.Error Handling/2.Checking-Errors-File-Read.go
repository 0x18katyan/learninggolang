package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("testfile.txtasd") // throws error

	//  f, err := os.Open("testfile.txt") // running code

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(string(bs))
}
