package main

import "fmt"

const LoginToken = "cat is lying"

func main() {
	var name string = "pranesh"
	fmt.Println(name)
	fmt.Printf("Type of this variable : %T \n ", name)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of this variable : %T \n ", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type of this variable : %T \n ", smallVal)

	var smallFloat float64 = 255.45367235452
	fmt.Println(smallFloat)
	fmt.Printf("Type of this variable : %T \n ", smallFloat)

	//default values
	var variableNumber int
	fmt.Println(variableNumber)
	fmt.Printf("Type of this variable : %T \n ", variableNumber)

	//implicit type
	var greet = "hello"
	fmt.Println(greet)
	fmt.Printf("Type of this variable : %T \n ", greet)

	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Printf("Type of this variable : %T \n ", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Type of this variable : %T \n ", LoginToken)
}
