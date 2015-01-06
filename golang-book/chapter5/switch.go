package main

import "fmt"

func main() {
	for i:= 0; i <= 9; i++{
		// Switch Case
		switch i {
		case 0: fmt.Println("zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		default: fmt.Printf("%d\n", i)
		}
	}
}
