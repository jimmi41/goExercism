package pangram
import "strings"

func IsPangram(s string) bool {
	 s = strings.ToLower(s)

    cnt := make([]int, 26)

    for _, ch := range s {

        if ch >= 'a' && ch <= 'z' {

            cnt[ch-'a']++
        }
    }

    for i := 0; i < 26; i++ {

        if cnt[i] == 0 {
            return false
        }
    }

    return true
}
