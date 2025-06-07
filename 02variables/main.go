package main

import "fmt"

//wtTOken:= 300000 not allowed out of the func

const LoggedinToken string = "jjjjjjdsdfdf12@#@" //public so initial letter should be capital like LoggedinToken

func main() {
	var username string = "shrey"
	fmt.Println(username)
	fmt.Printf("variable type: %T \n ", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("variable type: %T \n ", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable type: %T \n ", smallVal)

	var smallfloat float64 = 225.636464643561453646463
	fmt.Println(smallfloat)
	fmt.Printf("variable type: %T \n ", smallfloat)

	// impilicit way
	var WEBSITE = "IAMMANOFCULTURE.COM"
	fmt.Println(WEBSITE)
	fmt.Printf("variable type: %T \n ", WEBSITE)

	numberuser := 600000000
	fmt.Println(numberuser)

	fmt.Println(LoggedinToken)
	fmt.Printf("variable type: %T \n ", LoggedinToken)
}
