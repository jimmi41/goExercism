package differenceofsquares

func SquareOfSum(n int) int {
	var sum = 0
    for i := 0; i< n; i++{
        sum = sum + i+1
    }
    return sum*sum
}

func SumOfSquares(n int) int {
	var sum = 0
    for i := 0; i<n;i++{
        sum = sum + (i+1)*(i+1)
    }
    return sum
}

func Difference(n int) int {
	return SquareOfSum(n)-SumOfSquares(n)
}
