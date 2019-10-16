package main

import (
	"bufio"
	"fmt"
	"os"
)

type interestsStruct struct {
	primaryInterest   string
	secondaryInterest string
	tertiaryInterest  string
}
type personStruct struct {
	firstname, lastname string
	age                 int
	interests           interestsStuct
}

func main() {
	inputSlice := []personStruct{}

	m := map[int][]personStruct{}

	x := []string{"First Name", "Last Name", "Age", "Primary Interest", "Secondary Interest", "Tertiary Interest"}

	for k, v := range x {
		fmt.Printf("key %v, value %v of type %T\n", k, v, v)
		//inputSlice[k] = takeInput(v)
	}

}

func takeInput(input string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(fmt.Sprintf("Enter your %v: ", input))
	userInput, _ := reader.ReadString('\n')
	fmt.Println(userInput)
	return userInput
}
