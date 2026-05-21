package thefarm

import (
	"errors"
	"strconv"
)

// DivideFood calculates food per cow.
func DivideFood(fodder FodderCalculator, numOfCow int) (float64, error) {

	amount, err := fodder.FodderAmount(numOfCow)

	if err != nil {
		return 0, err
	}

	factor, err := fodder.FatteningFactor()

	if err != nil {
		return 0, err
	}

	foodPerCow := (amount * factor) / float64(numOfCow)

	return foodPerCow, nil
}

// ValidateInputAndDivideFood validates cows count before dividing food.

func ValidateInputAndDivideFood(
    fodder FodderCalculator,
    numOfCow int,
) (float64, error) {

    if numOfCow <= 0 {
        return 0, errors.New("invalid number of cows")
    }

    return DivideFood(fodder, numOfCow)
}
// ValidateNumberOfCows validates cows count.
func ValidateNumberOfCows(numOfCow int) error {

    if numOfCow < 0 {

        return errors.New(
            strconv.Itoa(numOfCow) +
                " cows are invalid: there are no negative cows",
        )
    }

    if numOfCow == 0 {

        return errors.New(
            "0 cows are invalid: no cows don't need food",
        )
    }

    return nil
}