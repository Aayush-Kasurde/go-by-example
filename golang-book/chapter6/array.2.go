package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 1
	x[1] = 2 
	x[2] = 3
	x[3] = 5
	var total float64
	for i := 0; i < 4; i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
	
	total = 0
	for _, value := range x {
		total += value 
	}
	
	fmt.Println(total / float64(len(x)))
	
	x = [5]float64{1,2,3,4,5}
	total = 0
	for _, value := range x {
		total += value 
	}
	
	fmt.Println(total / float64(len(x)))
}