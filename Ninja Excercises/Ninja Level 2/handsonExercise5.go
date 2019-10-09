package main

import (
	"fmt"
)

func main() {
	x := `This is a string which is made by using Raw String so I can input any fucking
  thing in here and it won't cause the compiler any issues like $&*!%(#&@!)(#)_)+(!~@)
  or %T%V%ASD%~5ASdN:./,
  All the above will only be taken as string and nothing else. Isn't it awesome?
  func main(){
    Just Showing Off Now!
    }`
	fmt.Println(x)
}
