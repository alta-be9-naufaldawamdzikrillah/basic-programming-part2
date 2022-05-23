package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FaktorBilangan(n int) string {
	// your code here
	var tampung []string

	for i := 1; i <= n; i++ {
		if n%1 == 0 {
			tampung = append(tampung, strconv.Itoa(i))
		}
	}

	return strings.Join(tampung, "\n")

}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilangan(number))
}
