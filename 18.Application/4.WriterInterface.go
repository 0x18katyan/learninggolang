package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s := "Hello World!\n"
	fmt.Fprint(os.Stdout, s)
	io.WriteString(os.Stdout, s)
}kb
