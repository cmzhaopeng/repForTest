package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func (p *Point) ScaleBy(factor int) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) ScaleBy2(factor int) Point {
	p.X *= factor
	p.Y *= factor
	return p
}

func main() {
	//test the difference between the pointer and the value receiver
	p := Point{1, 2}
	p.ScaleBy(2)
	fmt.Println(p) // {2, 4}
	p1 := Point{1, 2}
	p2 := p1.ScaleBy2(2)
	fmt.Println(p1) // {1, 2}
	fmt.Println(p2) // {2, 4}
}
