package largestseriesproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	// Convert to rune slice to handle multi-byte characters safely
	runeDigits := []rune(digits)

	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if len(runeDigits) < span {
		return 0, errors.New("span must be smaller than string length")
	}
	if span == 0 {
		return 1, nil
	}

	var maxProduct int64 = 0

	// Outer loop: Tracks the starting point of our span window
	for i := 0; i <= len(runeDigits)-span; i++ {
		var currentProduct int64 = 1

		// Inner loop: Iterates directly through the span window using indices
		for j := i; j < i+span; j++ {
			ch := runeDigits[j] // Grab the character directly by its index position

			if ch < '0' || ch > '9' {
				return 0, errors.New("digits string must only contain digits")
			}
			currentProduct *= int64(ch - '0')
		}

		if currentProduct > maxProduct {
			maxProduct = currentProduct
		}
	}

	return maxProduct, nil
}
