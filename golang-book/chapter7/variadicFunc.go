package main

import "fmt"

func main() {
	learnVariadicParams("Great", "Good", "Best");
}

func learnVariadicParams(myStrings ...interface{}) {
	for i, param := range myStrings {
		fmt.Println("Param [", i ,"] : ", param)
	}
	
	// Pass variadic value as a variable parameter
	
	fmt.Println("Params passed to this function :" , fmt.Sprintln(myStrings ...))
}