package lineup
import "strconv"
func Format(name string, n int) string {
    str := "th"
    if(n%10==1){
        if(n%100!=11){
            str = "st"
        }
    }else if(n%10==2){
        if(n%100!=12){
            str = "nd"
        }
    }else if(n%10==3){
        if(n%100!=13){
            str = "rd"
        }
    }
    return name+", you are the "+strconv.Itoa(n)+str+" customer we serve today. Thank you!"
}
