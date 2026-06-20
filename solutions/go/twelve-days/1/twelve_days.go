package twelvedays


// 1. Lists for the ordinal day names and the gifts for each day
var days = [13]string{
	"", "first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

var gifts = [13]string{
	"",
	"a Partridge in a Pear Tree.",
	"two Turtle Doves, ",
	"three French Hens, ",
	"four Calling Birds, ",
	"five Gold Rings, ",
	"six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ",
	"eight Maids-a-Milking, ",
	"nine Ladies Dancing, ",
	"ten Lords-a-Leaping, ",
	"eleven Pipers Piping, ",
	"twelve Drummers Drumming, ",
}

// Verse returns the exact string for a single verse (1 to 12).
func Verse(v int) string {
	// Start with the standard introduction line
	result := "On the " + days[v] + " day of Christmas my true love gave to me: "

	// Add the gifts in reverse order (from today down to day 1)
	for i := v; i > 0; i-- {
		// If we are on day 2 or higher, add "and " right before the very last gift
		if v > 1 && i == 1 {
			result += "and "
		}
		result += gifts[i]
	}

	return result
}

// Song returns all 12 verses joined together by newlines.
func Song() string {
	result := ""

	for i := 1; i <= 12; i++ {
		result += Verse(i)
		
		// Add a newline after every verse except the final 12th verse
		if i < 12 {
			result += "\n"
		}
	}

	return result
}
