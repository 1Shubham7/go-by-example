package main

import ("fmt")

func main() {
	x,y := sumAndMultiple(1,2)
	fmt.Print(x,y)
}

func sumAndMultiple(a,b int) (int, int) {
	return a+b, a*b
}