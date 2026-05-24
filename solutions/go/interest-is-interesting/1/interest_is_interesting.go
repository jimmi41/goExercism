package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if(balance < 0){
        return 3.213
    }
    if(balance < 1000){
        return 0.5
    }
    if(balance < 5000){
        return 1.621
    }
    return 2.475
    
}

func Interest(balance float64) float64 {
	if(balance == 0){
        return 0
    }
    return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	if(balance == 0){
        return 0
    }
    return balance + Interest(balance)
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {

    years := 0

    for balance < targetBalance {

        balance = AnnualBalanceUpdate(balance)

        years++
    }

    return years
}
