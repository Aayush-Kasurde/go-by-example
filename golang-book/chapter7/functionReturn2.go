package main

import "fmt"

func returnF() (int, int) {
	return 6, 4
}

func main() {
	a, b := returnF()
	fmt.Printf("A : %d, B : %d",a, b)
}
