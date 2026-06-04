package reversestring

func Reverse(s string) string {
	 r := []rune(s)

    left := 0
    right := len(r) - 1

    for left < right {

        r[left], r[right] =
            r[right], r[left]

        left++
        right--
    }

    return string(r)
}
