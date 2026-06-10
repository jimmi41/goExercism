package bottlesong


func Recite(startBottles, takeDown int) []string {

	upper := map[int]string{
		10: "Ten",
		9:  "Nine",
		8:  "Eight",
		7:  "Seven",
		6:  "Six",
		5:  "Five",
		4:  "Four",
		3:  "Three",
		2:  "Two",
		1:  "One",
	}

	lower := map[int]string{
		10: "ten",
		9:  "nine",
		8:  "eight",
		7:  "seven",
		6:  "six",
		5:  "five",
		4:  "four",
		3:  "three",
		2:  "two",
		1:  "one",
		0:  "no",
	}

	var song []string

	for bottles := startBottles; bottles > startBottles-takeDown; bottles-- {

		current := upper[bottles]
		next := lower[bottles-1]

		currentBottle := "bottles"
		if bottles == 1 {
			currentBottle = "bottle"
		}

		nextBottle := "bottles"
		if bottles-1 == 1 {
			nextBottle = "bottle"
		}

		song = append(song,
			current+" green "+currentBottle+" hanging on the wall,",
			current+" green "+currentBottle+" hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be "+next+" green "+nextBottle+" hanging on the wall.",
		)

		if bottles > startBottles-takeDown+1 {
			song = append(song, "")
		}
	}

	return song
}