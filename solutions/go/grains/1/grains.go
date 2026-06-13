package grains
import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("square must be between 1 and 64")
	}
    // 2. Calculate 2^(number-1) using bitwise left shift
	return 1 << (number - 1), nil
}

// Total calculates the total number of grains on all 64 squares.
func Total() uint64 {
	var count uint64 = 0
	
	// 3. Loop through all 64 squares (1 to 64)
	for i := 1; i <= 64; i++ {
		// 4. Properly unpack both the value and error variables
		val, _ := Square(i) 
		count += val
	}
	
	return count
}
