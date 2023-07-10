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

// closures function : here the intSeq function returns another function
// that function returns i, therefore whenever you call intSeq, another function 
// is created. Look at the code carefull, you will get it!

func intSeq() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}
