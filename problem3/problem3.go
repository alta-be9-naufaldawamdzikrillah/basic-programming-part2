package main

import (
	"fmt"
)

func PrimeNumber(number int) bool {
	// your code here
	if number >= 2 {
		for i := 2; i < number; i++ {
			if number%1 == 0 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func main() {
	// fmt.Println(PrimeNumber(0))
	fmt.Println(PrimeNumber(3))
	fmt.Println(PrimeNumber(7))
	fmt.Println(PrimeNumber(10))
}
