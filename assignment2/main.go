package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{
		height: 24.0,
		base:   12.0,
	}

	sq := square{
		sideLength: 32.0,
	}

	printArea(t)
	printArea(sq)
}
