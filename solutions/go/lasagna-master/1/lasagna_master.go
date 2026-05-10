package lasagnamaster
// import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {

    // default value
    if avgTime == 0 {
        avgTime = 2
    }

    return len(layers) * avgTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {

    noodles := 0
    sauce := 0.0

    for i := 0; i < len(layers); i++ {

        if layers[i] == "noodles" {
            noodles += 50
        }

        if layers[i] == "sauce" {
            sauce += 0.2
        }
    }

    return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {

    // last item from friend's list
    secret := friendsList[len(friendsList)-1]

    // replace last item in my list
    myList[len(myList)-1] = secret
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {

    // new slice
    result := make([]float64, len(quantities))

    // scaling factor
    factor := float64(portions) / 2

    // scale each item
    for i := 0; i < len(quantities); i++ {
        result[i] = quantities[i] * factor
    }

    return result
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
