package main

import "fmt"

func main() {
	for i:= 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("Fizz-Buzz")
		} else if i % 3 == 0 {
			fmt.Printf("Fizz ")
		} else if i % 5 == 0 {
			fmt.Printf("Buzz ")
		} else { 
			fmt.Printf("%d ", i)
		}
		
		if i % 10 == 0{
			fmt.Println("")
		}
	}
}