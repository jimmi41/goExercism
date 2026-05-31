package microblog

func Truncate(text string) string {
   // 1. Convert the string to a slice of runes (characters)
    runes := []rune(text)
    
    // 2. If the string is 5 characters or less, return it as-is
    if len(runes) <= 5 {
        return text
    }
    
    // 3. Slice the first 5 characters and convert back to a string
    return string(runes[0:5])
}
