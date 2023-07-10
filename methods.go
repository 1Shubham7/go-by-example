package main

import (
	"fmt"
)

type circle struct{
	radius int
}

func (r circle) area() (int) {
	return r.radius * 2 *3
}

func main() {
	a := circle{
		10,
	}

	fmt.Print(a.area())
}