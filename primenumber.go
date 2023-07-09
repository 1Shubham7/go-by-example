package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number :")
	var number int
	fmt.Scanln(&number)
	fmt.Println(isPrime(number))

}

func isPrime(number int) bool {
	var prime int
	for i:=2; i<number; i++ {
		if (number/i == 0){
			prime = 0
			break
		} 
	}
}