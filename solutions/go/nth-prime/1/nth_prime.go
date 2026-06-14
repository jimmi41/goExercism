package nthprime
import "errors"
// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
    if(n<1){
        return 0,errors.New("n must be greater than 0") 
    }
    
	if n == 1 {
		return 2, nil
	}

	
	count := 1       
	candidate := 3   

	
	for count < n {
		if isPrime(candidate) {
			count++
			if count == n {
				return candidate, nil
			}
		}
		candidate += 2
	}

	return candidate, nil
}

func isPrime(num int) bool {
	// Since we skip even numbers in Nth(), we only need to test odd factors here.
	// We stop checking once i * i > num because a composite number 
	// must have a factor less than or equal to its square root.
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
