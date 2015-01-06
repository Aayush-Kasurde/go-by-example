package main

import "fmt"

func main() {
	// Notify user for entering value 
	fmt.Print("Enter your name ")
	var input string 
	// Take value from user and save it input variable
	fmt.Scanf("%s", &input)
	output := input
	// Print value
	fmt.Printf("Hello %s", output)
}
