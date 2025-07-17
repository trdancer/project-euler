package probfive

/*
 2520 is the smallest number that can be divided by each of the numbers from to

without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from
1 to 20?
*/

func SmallestMultiple(n int) int {
	if n == 1 {
		return n
	}
	prevSmallest := SmallestMultiple(n - 1)
	i := 1
	for {
		prod := n * i
		if prod%prevSmallest == 0 {
			return prod
		}
		i++
	}
}
