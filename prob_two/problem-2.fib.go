package prob_two

func FibEvenSum(max int) int {
	sum := 0
	fib := [2]int{1, 2}
	for fib[0] <= max {
		if fib[0]%2 == 0 {
			sum += fib[0]
		}
		next := fib[0] + fib[1]
		fib[0] = fib[1]
		fib[1] = next
	}
	return sum
}
