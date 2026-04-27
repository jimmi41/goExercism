package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (successRate / 100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(float64(productionRate) * (successRate / 100) / 60)
}

func CalculateCost(carsCount int) uint {
    tens := carsCount / 10
    remainder := carsCount % 10

    return uint(tens*95000 + remainder*10000)
}