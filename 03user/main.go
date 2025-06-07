package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to the user input programme "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter yours rating on this ")

	//comma ok //err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating,", input)
	fmt.Printf("type of varialbe is %T", input)
}
