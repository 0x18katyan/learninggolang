package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("The OS of this System is:", runtime.GOOS)
	fmt.Println("The ARCH of this System is:", runtime.GOARCH)
}
