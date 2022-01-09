package main

import (
	"fmt"
	"strconv"
)

func Prima(N int) bool {
	// your code here
}

func FullPrima(n int) bool {
	digitPrima := map[int]bool{
		2: true,
		3: true,
		5: true,
		7: true,
	}
	if !Prima(n) {
		return false
	}
	str := strconv.Itoa(n)
	for _, s := range str {
		val, _ := strconv.Atoi(string(s))
		if _, exist := digitPrima[val]; !exist {
			return false
		}
	}
	return true // write your code
}

func main() {
	fmt.Println(FullPrima(2))  // true
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true
}
