package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	
	if len(rhyme) == 0 {
		return nil
	}

	
	var ans []string

	
	for i := 0; i < len(rhyme)-1; i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		ans = append(ans, line)
	}

	closingLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	ans = append(ans, closingLine)

	return ans
}
