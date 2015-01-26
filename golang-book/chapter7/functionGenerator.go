package main

import "fmt"

func makeGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i 
		i += 2
		return
	}
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	fmt.Println("Even Number generator")
	nextEven := makeGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
	
	fmt.Println("Odd number generator")
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 3
	fmt.Println(nextOdd()) // 5
}
