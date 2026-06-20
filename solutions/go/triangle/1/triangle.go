package triangle

// 1. We create 'Kind' as our own version of an integer.
type Kind int

// 2. We manually assign plain numbers to our categories.
const (
	NaT Kind = 0 // Not a triangle
	Equ Kind = 1 // Equilateral
	Iso Kind = 2 // Isosceles
	Sca Kind = 3 // Scalene
)

// KindFromSides checks the side lengths and returns the type.
func KindFromSides(a float64, b float64, c float64) Kind {
	
	// 3. Reject zero or negative side lengths
	if a <= 0 {
		return NaT
	}
	if b <= 0 {
		return NaT
	}
	if c <= 0 {
		return NaT
	}

	// 4. Reject if the triangle inequality rule is broken
	if a + b < c {
		return NaT
	}
	if a + c < b {
		return NaT
	}
	if b + c < a {
		return NaT
	}

	// 5. Check if all sides are equal
	if a == b {
		if b == c {
			return Equ
		}
	}

	// 6. Check if any two sides are equal
	if a == b {
		return Iso
	}
	if b == c {
		return Iso
	}
	if a == c {
		return Iso
	}

	// 7. If it passed all checks above, it must be scalene
	return Sca
}
