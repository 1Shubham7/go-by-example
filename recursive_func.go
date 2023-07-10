package main

import (
	"fmt"
)

func factorial(a int) int{
	if a<=1 {
		return 1
	}
		return a*factorial(a-1)
}

func main() {
	fmt.Print(factorial(5))
}