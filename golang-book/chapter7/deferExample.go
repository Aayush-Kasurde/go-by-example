package main

import "fmt"

func func1() {
	fmt.Println("I am function one")
}

func func2() {
	fmt.Println("I am function two")
}


func main() {
	defer func2()
	func1()
}
