package main

import (
	"fmt"
)

func main() {
	a := 10
	var b *int = &a
	*b = 100
	fmt.Print(a)
}