package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Printf("1 + 2 = %d", add(1,2))
}
