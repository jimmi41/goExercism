package atbashcipher

import "strings"

func Atbash(s string) string {

    s = strings.ToLower(s)

    result := ""

    count := 0

    for _, ch := range s {

        var out rune

        if ch >= 'a' && ch <= 'z' {

            out = 'z' - (ch - 'a')

        } else if ch >= '0' && ch <= '9' {

            out = ch

        } else {

            continue
        }

        if count > 0 && count%5 == 0 {

            result += " "
        }

        result += string(out)

        count++
    }

    return result
}