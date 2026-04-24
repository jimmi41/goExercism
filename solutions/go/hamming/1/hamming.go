package hamming
import "errors"

func Distance(a, b string) (int, error) {
	var la = len(a)
    var lb = len(b)

    if(la!=lb){
        return 0,errors.New("length mismatch")
    }
 
	count := 0
	for i := 0; i < len(a); i++ {
    	if a[i] != b[i] {
        	count++
    	}
	}
    return count,nil
}
