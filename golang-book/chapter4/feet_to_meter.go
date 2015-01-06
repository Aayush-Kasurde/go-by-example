package main

import "fmt"

func main() {
	// Define constant for conversion
	const x float32 = 0.3048
	// Notify user to input feet value
	fmt.Print("Enter feet to convert to meter  :")
	var input float32
	// Take user input
	fmt.Scanf("%f", &input)
	// Calculate output
	output := input * x
	// Print Output
	fmt.Printf("%.4ff = %.4fm", input, output)
}
