package sumofmultiples


func SumMultiples(limit int, divisors ...int) int {
    sum := 0

    for i := 1; i < limit; i++ {

        for j := 0; j < len(divisors); j++ {

            d := divisors[j]

            if d != 0 && i%d == 0 {
                sum += i
                break
            }
        }
    }

    return sum
}