package main

import (
	"fmt"
)

func main() {
	fmt.Print(sume(1,2,3,4,5))
}

//variadic function (slice... int)

func sume(nums ... int) int{
	total := 0
	for _, num := range nums{
		total += num
	}
	return total
}
