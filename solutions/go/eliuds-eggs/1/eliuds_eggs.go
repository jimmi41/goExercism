package eliudseggs

func EggCount(displayValue int) int {
    if(displayValue <= 0){
        return 0
    }
    count := 0
    for displayValue > 0{
        if(displayValue & 1 == 1){
            count++
        }
        displayValue = displayValue>>1
    }
    return count
}
