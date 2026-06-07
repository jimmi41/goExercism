package armstrongnumbers
import "math"

func IsNumber(num int) bool {
    if num < 0 {
		return false
	}

	// Step 1: Count the total number of digits
	temp := num
	count := 0
	for temp > 0 {
		count++
		temp /= 10
	}

	// Step 2: Extract each digit, raise to power, and sum them up
	temp = num
	sum := 0
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(count)))
		temp /= 10
	}

	// Step 3: Compare total sum to original number
	return sum == num
}
