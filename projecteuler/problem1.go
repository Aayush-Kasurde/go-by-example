// Problem 1 - http://projecteuler.net
//
// Add all the natural numbers below 1000 that 
// are multiples of 3 or 5

package main

import "fmt"

func main() {
	sum := 1
	i := 1
	for i <= 1000 {
		if i % 3 == 0 || i % 5 == 0 {
			sum = sum + i
		}
		i = i + 1
	}
	fmt.Printf("Sum of all number below 1000 which are divisible by 3 or 5 : %d", sum)
}