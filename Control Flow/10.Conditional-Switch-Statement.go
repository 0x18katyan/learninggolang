package main

import (
	"fmt"
)

func main() {
	fmt.Println("switch in func main()")
	switch {
	case false:
		fmt.Println("this should not print")
	case true:
		fmt.Println("this should print")
	case 2 == 2:
		fmt.Println("this should print")
	}
	othermain()
	stepmain()
}

/* Fall through is when one case has been satisfied
then the other case need not be satisfied to be executed
*/

func othermain() {
	fmt.Println("switch in func othermain()")
	switch {
	case true:
		fmt.Println("This should be printed because the case is true.")
		fallthrough
	case false:
		fmt.Println(`This wouldn't have been printed if the fallthrough statement wasn't present in the above code.`)
		//fallthrough
	case 2 == 2:
		fmt.Println("\n")
		fmt.Println(`This should not print because fallthrough will only execute the block after it's case
without checking the case after it at all. If this is to be printed then another fallthrough needs to be added to the previous case
which should necessarily be if this is printed in the screen.`)
	}
}

/* When using case, only a single case can be matched in a switch without the use of fallthrough
in that sense it is not like 'if and else if' where if something is not matched another condition can be checked.
Thus, when using case, only highly specific case is generally used.

NOTE: A block inside a case will only be executed if the case is true and other types like int cannot be used,
i.e for declaring a case of even or odd, a boolean value needs to be assigned to a variable then checked through
case whether the boolean is true or not.

For example see the following function:
*/

func stepmain() {
	fmt.Println("switch in stepmain()")
	for x := 0; x < 10; x++ {
		y := (x%2 == 0)
		fmt.Printf("y is now %v thus ", y)
		switch {
		// NOTE: if 'switch y {' is used in the above line then only the true cases will be printed.
		/*
			case x == 1:
				//Only a single case is checked through an iteration.
				fmt.Printf("Value is %v.\n", x)
		*/
		case y == true:
			//Only this block will be ever executed.
			fmt.Printf("%v is even.\n", x)
		case y == false:
			fmt.Printf("%v is odd.\n", x)
			// Unless the fallthrough below is used.
			fallthrough
		case x == 1:
			fmt.Printf("%v is odd therefore this is being printed.\n", x)
		}
	}
}
