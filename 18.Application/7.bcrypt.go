package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `password123`
	fmt.Println(s)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(bs)

	toCheckPassword := "wrongPassword"

	err = bcrypt.CompareHashAndPassword(bs, []byte(toCheckPassword))

	if err != nil {
		fmt.Println("You can't Login")
		fmt.Println(err)
		return
	}

	fmt.Println("You're Logged In")

}
