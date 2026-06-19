package squareroot

func SquareRoot(n int) (int, error) {
	if(n==1){
        return 1,nil
    }
    l:=0
    h:=(n/2)+1
    ans := 0
    for l<=h{
        m := l + (h-l)/2
        if(m == n/m){
            return m,nil
        }else if(m > n/m){
            h=m-1
        }else{
            ans = m
            l=m+1
        } 
    }
    return ans,nil
}
