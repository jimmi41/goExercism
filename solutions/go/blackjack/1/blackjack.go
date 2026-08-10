package blackjack

func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	// split aces
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	player := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)

	// blackjack
	if player == 21 {
		if dealer == 10 || dealer == 11 {
			return "S"
		}
		return "W"
	}

	// stand
	if player >= 17 {
		return "S"
	}

	// 12-16
	if player >= 12 && player <= 16 {
		if dealer >= 7 {
			return "H"
		}
		return "S"
	}

	// <= 11
	return "H"
}