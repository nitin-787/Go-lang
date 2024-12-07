package main

import "fmt"


// capital letter in LoginToken - L means this var is treated as public var
const LoginToken string = "foobarbaz"

func main() {
	var username string = "nitin"
	fmt.Println(username)
	fmt.Printf("varibales is type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("varibales is type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type: %T \n", smallVal)

	var smallFloat float64 = 255.545353453495834
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type: %T \n", smallFloat)

	// default values
	var anotherString string
	fmt.Println(anotherString) // empty string

	var anotherInt int
	fmt.Println(anotherInt) // 0

	// implicit type
	var website = "gihtub.com/nitin-787" // once assigned, then later cannot change the type 
	fmt.Println(website)

	// no var style
	numberOfUser := 404
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
