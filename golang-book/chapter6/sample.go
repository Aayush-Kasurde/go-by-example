package main

import "fmt"

func main() {
	var x = []int{0,1,2,3,0,4,0,6,0}
	fmt.Println("Before bringing zero at start of array :", x)
	zerocount := 0
	for index, value := range x {
		if value == 0 {
			// if found zero then swap space of zero with end of zero at start 
			x[zerocount], x[index] = x[index], x[zerocount]
			zerocount++
		}
	}
	fmt.Println("After bringing zero at start of array  :", x)
}
