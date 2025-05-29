package main

import (
	"fmt"
)

func Multiples(n int) int {
	var sum int = 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	answer := Multiples(1000)
	fmt.Println(answer)
}
