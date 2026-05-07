package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	var sol []string
    for i:=0;i<len(candidates);i++{
        if(checkA(subject,candidates[i])){
            sol = append(sol,candidates[i])
        }
    }
    return sol
}

func checkA(a, b string) bool {

    if len(a) != len(b) {
        return false
    }
    a = strings.ToLower(a)
	b = strings.ToLower(b)
    if a == b {
    return false
}
    count := make(map[rune]int)
    
    for _, ch := range a {
        count[ch]++
    }

    for _, ch := range b {
        count[ch]--
    }

    for _, value := range count {
        if value != 0 {
            return false
        }
    }
    return true
}
