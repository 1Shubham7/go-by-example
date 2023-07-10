package main

import (
	"fmt"
)

// company := struct{
// 	name string
// 	code int
// } {
// 	"LG",
// 	100,
// }

func main(){
	a := struct {
		name string
		code int
	} {
		"sss",
		100,
	}
	fmt.Print(a)
	// fmt.Print(company)
}