package main

import (
	"fmt"
	"time"
)

func main() {
	if (time.Now().Hour()<12){
		fmt.Print("Good Morning")
	} else {
		fmt.Print("Good Night")
	}
}