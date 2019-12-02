package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("testfile.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := strings.NewReader("WASSUP DAWG\n")

	io.Copy(f, r)
	othermain()
}
