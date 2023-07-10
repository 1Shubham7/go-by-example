package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("How many number you want to enter")
	fmt.Scan(&n)
	fmt.Print("Enter the numbers")
	arr := make([]int,n)
	for i:=0; i<n; i++ {
		fmt.Scan(&arr[i])
	}
	
	sum := 0
	for _,num := range arr {
		sum += num
	}
	fmt.Print(sum)
}