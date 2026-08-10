package luhn

func Valid(id string) bool {
    sum := 0
    double := false
    digitCount := 0

    // iterate from right to left
    for i := len(id) - 1; i >= 0; i-- {
        ch := id[i]

        // skip spaces
        if ch == ' ' {
            continue
        }

        // invalid character
        if ch < '0' || ch > '9' {
            return false
        }

        digit := int(ch - '0')
        digitCount++

        if double {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }

        sum += digit
        double = !double
    }

    // must have at least 2 digits
    if digitCount <= 1 {
        return false
    }

    return sum%10 == 0
}