package main

import "fmt"

func main() {
	fmt.Println("welcome to the pointer ") //direct refrence to memory address give the actusl memory

	//var ptr *int //just to create the pointer *
	//fmt.Println("value of pointer ", ptr)

	mynumber := 23

	var ptr = &mynumber //& is used to create as well as refrencing//&

	fmt.Println("value of actuall ptr is ", ptr)  //pointer
	fmt.Println("value of actuall ptr is ", *ptr) //actual value of pointer

	*ptr = *ptr + 2
	fmt.Println("value of ppointer ", mynumber)
}
