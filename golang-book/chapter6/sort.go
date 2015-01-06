package main

import "fmt"

func main() {
	x := []int64 {12,52,3,5,6}
	for itemcount := len(x) - 1; ; itemcount-- {
		flag := false
		for index := 0; index < itemcount ; index++ {
			if x[index] > x[index + 1] {
				x[index], x[index + 1] = x[index + 1], x[index]
				flag = true
			}
		}
		
		if flag == false {
			break
		}
	}
	fmt.Println(x)
	x = []int64{12,3,2,34}
	for i := len(x)-1; i > 0; i-- {
		fmt.Println(i)
		for j := 0; j < i; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	fmt.Println(x)
	/*
	for i:= 0; i < len(x); i++ {
		for j, _ := range x {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	fmt.Println(x)
	*/
}