package isogram
import "strings"
func IsIsogram(word string) bool {
    seen := make(map[rune]bool)

    word = strings.ToLower(word)

    for _, ch := range word {

        if ch == ' ' || ch == '-' {
            continue
        }

        if seen[ch] {
            return false
        }

        seen[ch] = true
    }

    return true
}
