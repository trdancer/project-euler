package prob_four

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is

// 9009 = 91 * 99

// Find the largest palindrome made from the product of two
// 3-digit numbers.
/**
	// The largest number that can be made from the product of two N-digit numbers is
	// 9_n ^2
	// the smallest is 10_n^2
	// so for N=3 the answer A lies within 100^2 < A < 999^2 = 10,000 < A < 998,001
	// the answer cannot be prime because then its only factors would be 1 and itself, and since 1 is not a N digit number (for N > 1)

	// the closest palindrome number to 998,001 is 997,799
	=
				  9
			   90
			  700
		  7,000
		 90,000
	+ 900,000

	996699
	995599
	994499
	993399
	992299
	991199
	990099
	989989
	...
	11211
	11111
	11011
	...
	10301
	10201
	10101
	10001

	9 * 1 = 9
	= (9 * 1) * (1)
	91 * 99 =
	(1 * 99) + (90 * 99) =
  (1 * ((9 * 1) + (9 * 10))) + ((9 * 10) * ((9 * 1) + (9 * 10))) =
	((9 * 1) + (9 * 10)) + ((9 * 10) * (9 * 1)) + ((9 * 10) * (9 * 10)) =
	((9 * 1) + (9 * 10)) + (9 * 9 * 10) + (9 * 9 * 10 * 10) =
	((9^1 * 10^0) + (9^1 * 10^1)) + (9^2 * 10^1) + (9^2 * 10^2)

	LHS = (9 * 1) + (9 * 10) + (9 * 100) = 999
	RHS =
*/
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
				fmt.Printf("%v is a palidrome number made by factors %v and %v\n", product, i, j)
				break
			}
		}
	}
	return answer

}
