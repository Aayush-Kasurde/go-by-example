package main

import "fmt"

type Point struct {
	x, y int;
};

func (p Point) Print() {
	fmt.Printf("(x = %d, y = %d)", p.x, p.y)
}

func main() {
	p := Point{10, 20};
	fmt.Printf("%v", p)
	p.Print()
}