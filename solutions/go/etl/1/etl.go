package etl
import "strings"

func Transform(in map[int][]string) map[string]int {
    var newMap = make(map[string]int)
    
    for key, value := range in {
        for _,str := range value{
            newMap[strings.ToLower(str)] = key
        }
    }
    return newMap
}
