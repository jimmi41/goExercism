package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int { 
    rMap := map[string]int{
        "black":0,
        "brown":1,
        "red":2,
        "orange":3,
        "yellow":4,
        "green":5,
        "blue":6,
        "violet":7,
        "grey":8,
        "white":9,
    }
    val := 0
    // for _,color := range(colors){
        val = val*10 + rMap[colors[0]]
    	val = val*10 + rMap[colors[1]]
    // }
    return val
}
