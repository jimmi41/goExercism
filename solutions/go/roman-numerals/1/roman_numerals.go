package romannumerals

import "errors"

type numeralMapping struct {
	value  int
	symbol string
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("input must be between 1 and 3999")
	}

	conversions := []numeralMapping{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""

	for _, mapping := range conversions {
		for input >= mapping.value {
			roman += mapping.symbol   
			input -= mapping.value    
		}
	}

	return roman, nil
}
