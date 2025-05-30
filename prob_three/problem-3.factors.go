package prob_three

/*
* Note that the largest factor of a number will always be multiplied by another
* factor which is the smallest factor of that number
* The non 1/identity factors of a number will only be between [2, n/2]
* so to find the largest factor we can start from 2 and find the first divisible number
* of N

156
/\
2 78

	  /\
	 2 39
	  	/\
		 3 13
*/
func getLargestFactor(n int) (int, bool) {
	end := n / 2
	for i := 2; i < end; i++ {
		// n is divisible by i
		if n%i == 0 {
			// get the cofactor of i
			cofactor := n / i
			// get the largest factor of the cofactor
			lfFactor, _ := getLargestFactor(cofactor)
			// cofactor is prime if the largest factor is itself
			// short circuit func here
			return cofactor, lfFactor == cofactor
		}
	}
	// No divisors, so N is largest factor of itself and is prime
	return n, true
}

func LargestPrimeFactor(n int) int {

	factor, prime := getLargestFactor(n)
	for {
		if prime {
			break
		}
		factor, prime = getLargestFactor(factor)
	}
	return factor
}
