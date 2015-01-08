package main

import "fmt"

func main() {
	add1 := func(x int) int {
		return x + 1
	}
	
	n := 7
	fmt.Println("n = ", n)
	n = add1(n)
	fmt.Println("Add1(n) returns ", n)
}