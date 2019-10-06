package main

import (
	"bufio"
	"fmt"
	"os"
)

type interests struct {
	primaryInterest   string
	secondaryInterest string
	tertiaryInterest  string
}
type personStruct struct {
	firstname, lastname string
	age                 int
	interests           interests
}

var person personStruct

func main() {
	person.firstname = takeInput("First Name")
	person.lastname = takeInput("Last Name")
	person.age = takeInput("Age")
	person.interests.primaryInterest = takeInput("Primary Interest")
	person.interests.secondaryInterest = takeInput("Secondary Interest")
	person.interests.tertiaryInterest = takeInput("Tertiary Interest")
}

func takeInput(input string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(fmt.Sprint("Enter your %d:  ", input))
	user_input, _ := reader.ReadString("\n")
	return user_input
}
