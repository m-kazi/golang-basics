package main

import "fmt"

// public token starts with capital letter, public can be accessed by any files in the folder
const LoginToken  string = "lkjnlijfwjbkjbs"

func main() {
	var username string = "Tanim"
	fmt.Println(username)
	fmt.Printf("Variable type is: %T \n", username)


	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is: %T \n", isLoggedIn)

	// in general we can use "int"
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable type is: %T \n", smallVal)


	var smallFloat float64 = 255.123466789375678
	fmt.Println(smallFloat)
	fmt.Printf("Variable type is: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is: %T \n", anotherVariable)


	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable type is: %T \n", anotherString)
	

	// implicit type
	var website = "kazzcode.com"
	fmt.Println(website)

	// no var style, only can be used inside a function not globally
	numberOne := 50000
	fmt.Println(numberOne)

	fmt.Println(LoginToken)
	fmt.Printf("Variable type is: %T \n", LoginToken)


}