package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	arr := make([]string, 10)
    arr[0] = "black"
    arr[1] = "brown"
	arr[2] = "red"
    arr[3] = "orange"
    arr[4] = "yellow"
    arr[5] = "green"
    arr[6] = "blue"
    arr[7] = "violet"
    arr[8] = "grey"
    arr[9] = "white"
    return arr
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
    arr := Colors()
	for i:=0;i<10;i++{
        if(arr[i]==color){
            return i
        }
    }
    return -1
}
