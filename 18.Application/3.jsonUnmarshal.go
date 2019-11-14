package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Age         int
}

func main() {
	p1 := person{
		First: "Ashim",
		Last:  "Mahara",
		Age:   24,
	}
	p2 := person{
		First: "Hobo",
		Last:  "Scientist",
		Age:   24,
	}

	people := []person{p1, p2}

	x, err := json.Marshal(people)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	type person struct {
		First, Last string
		Age         int
	}
	fmt.Println("x is a marshal/ed []person which is: ", x)

	fmt.Printf("The type of x is: %T\n", x)

	fmt.Println("After Converting to String:", string(x))

	var x2 []person

	err = json.Unmarshal(x, &x2)

	if err != nil {
		fmt.Println("Error while unmarshal:", err)
	}

	fmt.Println("After unmarshal: ", x2)

	foo()

}

func foo() {
	x := `[{"FirstName":"Ashim","LastName":"Mahara","Age":24},{"FirstName":"Hobo","LastName":"Scientist","Age"
:24}]`

	type person struct {
		First string `json:"FirstName"`
		Last  string `json:"LastName"`
		Age   int
	}

	var x2 []person

	bs := []byte(x)

	err := json.Unmarshal(bs, &x2)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(x2)

}
