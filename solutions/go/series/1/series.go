package series

func All(n int, s string) []string {
	if(len(s)<n || n<=0){
        return nil
    }
    var srr = []string{}
    for i:=0;i<=len(s)-n;i++{
        srr=append(srr,s[i:i+n])
    }
    return srr
}

func UnsafeFirst(n int, s string) string {
	if(len(s)<n || n<=0){
        return ""
    }
    return s[:n]
}
