package darts
import "math"

func Score(x, y float64) int {
	
    var loc = math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
    if loc > 10 {
        return 0
    }else if loc > 5{
        return 1
    }else if loc > 1{
        return 5
    }else{
        return 10
    }
}
