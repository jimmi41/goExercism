package wordcount

import (
	"strings"
	"unicode"
)


type Frequency map[string]int

func WordCount(phrase string) Frequency {
	result := make(Frequency)
    
	phrase = strings.ToLower(phrase)

	splitFunc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '\''
	}

	rawWords := strings.FieldsFunc(phrase, splitFunc)

	
	for _, word := range rawWords {
		
		cleanWord := strings.Trim(word, "'")

		if cleanWord != "" {
			result[cleanWord]++
		}
	}

	return result
}
