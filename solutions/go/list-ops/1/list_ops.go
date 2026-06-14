package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	cnt := initial // Fixed: Initialized with 'initial' parameter instead of 0
	for _, val := range s {
		cnt = fn(cnt, val)
	}
	return cnt
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	accumulator := initial

	// Loop backwards: start at the last index, stop after index 0
	for i := len(s) - 1; i >= 0; i-- {
		// Standard functional programming convention for Foldr:
		// The current value is the first argument, accumulator is the second.
		accumulator = fn(s[i], accumulator)
	}

	return accumulator
}

func (s IntList) Filter(fn func(int) bool) IntList {
	items := IntList{}

	for _, val := range s {
		if fn(val) {
			items = append(items, val)
		}
	}

	return items
}

func (s IntList) Length() int {
	count := 0
	for range s { // Fixed: Removed double blank identifier syntax error
		count++
	}
	return count
}

func (s IntList) Map(fn func(int) int) IntList {
	for i, val := range s {
		s[i] = fn(val)
	}
	return s
}

func (s IntList) Reverse() IntList {
	i := 0
	j := len(s) - 1
	// Fixed: Removed naked blank identifier syntax error and changed '!=' to '<'
	for i < j {
		t := s[i]
		s[i] = s[j]
		s[j] = t
		i++
		j--
	}
	return s
}

// Append takes a second IntList and combines it into the current in-place structure
func (s IntList) Append(lst IntList) IntList {
	// We use the variadic '...' operator to unpack lst elements directly into s
	s = append(s, lst...)
	return s
}

// Concat takes a slice of IntLists and flattens them sequentially into a single IntList
func (s IntList) Concat(lists []IntList) IntList {
	// Loop through each separate list container in the provided slice
	for _, currentList := range lists {
		// Leverage our existing Append method to merge the sublist elements
		s = s.Append(currentList)
	}
	return s
}
