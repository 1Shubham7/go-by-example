package main

import (
	"fmt"
)

func main() {
	a := intSeq()
	
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func intSeq() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}
