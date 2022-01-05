package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Check() int {
	if p.X > 0 && p.Y > 0 {
		return 1
	} else if p.X > 0 && p.Y < 0 {
		return 4
	} else if p.X < 0 && p.Y > 0 {
		return 2
	} else if p.X < 0 && p.Y < 0 {
		return 3
	} else {
		return -1
	}
}

func main() {
	point := Point{80, -5}
	fmt.Println("Kuadrannya ada di ", point.Check())
}
