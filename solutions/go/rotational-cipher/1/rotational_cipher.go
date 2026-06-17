package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
    sol := ""
    shiftKey = shiftKey%26
    if(shiftKey == 0){
        return plain
    }
    for _, ch := range plain {
		switch {
		case ch >= 'a' && ch <= 'z':
			shiftedCh := 'a' + (ch-'a'+rune(shiftKey))%26
			sol = sol + string(shiftedCh)

		case ch >= 'A' && ch <= 'Z':
			shiftedCh := 'A' + (ch-'A'+rune(shiftKey))%26
			sol = sol + string(shiftedCh)

		default:
			sol = sol + string(ch)
		}
	}
    return sol
}
