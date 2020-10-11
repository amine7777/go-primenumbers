package main

import (
	"fmt"
)
func main () {
	fmt.Println("Please enter a number: ")
	var number int
	fmt.Scanln(&number)
	if number > 1 {
		for i := 2; i <= number; i++ {
			if number%2 == 0 {
				fmt.Println("This is a prime number")
				break
			
			} else {
				fmt.Println("This is not a prime number")
				break
			 }
		}
	
	} else {
		fmt.Println("This is not a prime number")
	}
}

