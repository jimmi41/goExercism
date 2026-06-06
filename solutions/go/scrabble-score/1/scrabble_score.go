package scrabblescore
import "strings"
func Score(word string) int {
    score := 0
    word = strings.ToUpper(word)
    for _,ch:=range word{
        if(ch=='A' || ch=='E' ||ch=='I' ||ch=='O' ||ch=='U' ||ch=='L' ||ch=='N' ||ch=='R' ||ch=='S' ||ch=='T'){
            score +=1
        }else if(ch=='D' || ch=='G'){
            score += 2
        }else if(ch=='B' || ch=='C' || ch=='M' || ch=='P'){
            score += 3
        }else if(ch=='F' || ch=='H' || ch=='V' || ch=='W'|| ch=='Y'){
            score += 4
        }else if(ch=='K'){
            score += 5
        }else if(ch=='J'||ch=='X'){
            score += 8
        }else if(ch=='Q'||ch=='Z'){
            score += 10
        }
    }
    return score
}
