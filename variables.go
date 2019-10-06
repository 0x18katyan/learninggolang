package main

import (
	"fmt"
)

//I am now going to declare a variable outside of a function which will be availabe inside all functions

var globalVarX = "My name is globalVarX, it's nice to meet you."

func main() {
	initialize_variables()
	get_Type()
	test_global()
}

func initialize_variables(){
	var x string
	// String Literals are used for the exact purpose as their names; between them lies anything that may come and they will be known as strings and nothing else
	x = `Hi, I am using string literal for this purpose @!(*&@*&^&*#$_!\|\{}:":>??><>?<?><?><@#`
	fmt.Println(x)

	var y = `This can be
	like this				or 				like this		
	but it's still the same for the compiler within the string literal`
	fmt.Println(y)
}

func get_Type() {
	var x string
	fmt.Printf("The variable x is of type %T\n", x)
	fmt.Println("I am now going to print a global variable, behold.")
	fmt.Println("\n")
	fmt.Println(globalVarX)
	fmt.Println("\n")
	fmt.Println(`And, I must not be able to change that variable's value with this statement which is :
	globalVarX = "It didn't work"`)
	fmt.Println("Let's check if it works or not (TBH I don't know about it as of now but I am going to find it out.")
	globalVarX = "It didn't work"
	fmt.Println(globalVarX)
}

func test_global() {
	fmt.Println("\n")
	fmt.Println(`So yeah, it didn't work.
	And this is a function named test_global(), hello!
	I am printing this with that function.
	And yes, with a string literal so all this is actually a single string.
	
	So, what I am now going to test is, after I assigned the variable globalVarX another value, I am going to check if that value
	is the changed version for other functions or not.
	I am going to do that with fmt.Println(globalVarX.`)

	if globalVarX == "It didn't work" {
		fmt.Println(globalVarX)

	} else {
		fmt.Println("The value didn't change!")
	}
}