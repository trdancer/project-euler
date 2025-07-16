package prob_four

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is

// 9009 = 91 * 99

// Find the largest palindrome made from the product of two
// 3-digit numbers.

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(n int) bool {
	asString := strconv.FormatInt(int64(n), 10)
	for i, c := range asString {
		z := len(asString) - i - 1
		opposite := asString[z]
		if c != rune(opposite) {
			return false
		}
	}
	return true
}

func LargestPalindrome(n int) int {
	MAX, err := strconv.Atoi(strings.Repeat("9", n))
	MIN, e := strconv.Atoi(fmt.Sprintf("1%v", strings.Repeat("0", n-1)))
	if err != nil || e != nil {
		return -1
	}
	answer := -1
	for i := MAX; i >= MIN; i-- {
		for j := i; j >= MIN; j-- {
			product := int(i * j)
			if isPalindrome(product) && product > answer {
				answer = product
				break
			}
		}
	}
	return answer

}
