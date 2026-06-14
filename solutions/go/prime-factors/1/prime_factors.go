package primefactors

func Factors(n int64) []int64 {
	// 1. Initialize the slice (pre-allocating space for efficiency)
	prime := make([]int64, 0, 64)
	
	// A placeholder divisor starting at the smallest prime number
	var divisor int64 = 2

	// 2. Loop to find factors
	for n > 1 {
		if n % divisor == 0 {
			prime = append(prime, divisor)
			
			n = n / divisor 
		} else {
			divisor++ 
		}
	}
	
	return prime
}
