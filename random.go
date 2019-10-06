package main

import (
	"fmt"
	"bufio"
	"os"
)

type interests_struct struct {
	primaryInterest string
	secondaryInterest string
	tertiaryInterest string
}

type person_type struct {
	firstName, lastName string
	age int
	interests interests_struct
}

func main() {
	var person person_type
	person.firstName = getInput(&"First Name")
	person.lastName = getInput("Last Name")
	person.age = getInput("Age")
	person.interests.primaryInterest = getInput("Primary Interest")
	person.interests.secondaryInterest = getInput("Secondary Interest")
	person.interests.tertiaryInterest = getInput("Tertiary Interest")
	fmt.Println(person)
}

func getInput(input string) {
	Reader := bufio.newReader(os.Stdin)
	toPrint := fmt.Sprintf("What is your %d?", input)
	fmt.Println(toPrint)
	saved_variable, _ := Reader.readString("\n")
	return saved_variable
}