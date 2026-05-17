package logs

func Application(log string) string {

    for _, ch := range log {

        if ch == '\u2757' {
            return "recommendation"

        } else if ch == '\U0001F50D' {
            return "search"

        } else if ch == '\u2600' {
            return "weather"
        }
    }

    return "default"
} 

func Replace(log string, oldRune, newRune rune) string {

    result := ""

    for _, ch := range log {

        if ch == oldRune {
            result += string(newRune)
        } else {
            result += string(ch)
        }
    }

    return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool { 
     count := 0

    for range log {
        count++
    }

    return count <= limit
}
