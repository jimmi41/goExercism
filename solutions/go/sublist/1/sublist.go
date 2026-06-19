package sublist

import "slices" // Import the core slices package

func Sublist(l1, l2 []int) Relation {
	// 1. If both are exactly identical (handles double empty list and identical lists)
	if slices.Equal(l1, l2) {
		return RelationEqual
	}

	// 2. If l1 is empty but l2 isn't, l1 is automatically a sublist
	if len(l1) == 0 {
		return RelationSublist
	}

	// 3. If l2 is empty but l1 isn't, l1 is automatically a superlist
	if len(l2) == 0 {
		return RelationSuperlist
	}

	// 4. If l1 is shorter, check if l2 contains all of l1
	if len(l1) < len(l2) {
		if contains(l2, l1) {
			return RelationSublist
		}
	}

	// 5. If l1 is longer, check if l1 contains all of l2
	if len(l1) > len(l2) {
		if contains(l1, l2) {
			return RelationSuperlist
		}
	}

	// 6. If none of the conditions match, they are completely mismatched
	return RelationUnequal
}

// contains is a helper function that checks if 'short' list is entirely contained inside 'long' list
func contains(long, short []int) bool {
	// Slide a window across the long list, stopping where the short list would overflow
	for i := 0; i <= len(long)-len(short); i++ {
		// Slice out a window of the long list and compare it directly to the short list
		if slices.Equal(long[i:i+len(short)], short) {
			return true
		}
	}
	return false
}
