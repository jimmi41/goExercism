package sieve

func Sieve(limit int) []int {
	if(limit<2){
        return nil
    }
    prime := make([]bool,limit+1)
    for i:=2;i*i<=limit;i++{
        if(prime[i]==true){
            continue
        }
        temp := i*i
        
        for temp <= limit{
            prime[temp] = true
            temp = temp + i
        }
    }
    var primes []int
	for i := 2; i <= limit; i++ {
		if !prime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
