package main

import "fmt"

func greatest(args []int) int {
	high := 0
	for _, v := range args {
		if v > high {
			high = v
		}
	}
	return high
}

func main() {
	xs := []int {10, 11, 21, 31}
	fmt.Println("Highest number is list : ", greatest(xs))
}
