package main

import "fmt"

func main() {
	// String manipulations
	var x string = "Hello World"
	fmt.Println(x)
	
	var y string
	y = "Hello World"
	fmt.Println(y)
	
	var z string 
	z = "Abhijeet"
	fmt.Println(z)
	z = "Anjali"
	fmt.Println(z)
	
	var i string 
	i = "first "
	fmt.Println(i)
	i = i + "Second"
	fmt.Println(i)
	i += " Third"
	fmt.Println(i)
	
	var s, t string 
	s = "Hello"
	t = "hello"
	
	fmt.Println(s == t)
	
	u := 1
	fmt.Println(u)
	
}
