package isbnverifier

func IsValidISBN(isbn string) bool {
    sum := 0
    digit := 10
    count := 0

    for i := 0; i < len(isbn); i++ {
        ch := isbn[i]

        if ch == '-' {
            continue
        }

        count++

        if count > 10 {
            return false
        }

        var value int

        if ch == 'X' {
            if count != 10 {
                return false
            }
            value = 10
        } else if ch >= '0' && ch <= '9' {
            value = int(ch - '0')
        } else {
            return false
        }

        sum += value * digit
        digit--
    }

    if count != 10 {
        return false
    }

    return sum%11 == 0
}