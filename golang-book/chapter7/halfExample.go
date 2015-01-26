package main

import "fmt"

func half(x int) (int, bool) {
	if x % 2 == 0 {
		return x / 2, true
	} else {
		return x, false
	}
}

func main() {
	fmt.Println(half(2))
	fmt.Println(half(3))
}

