package piscine

func Capitalize(s string) string {
	myarray := []rune(s)
	isString := true

	for i := 0; i < len(s); i++ {
		if checker(myarray[i]) && isString {
			if myarray[i] >= 'a' && myarray[i] <= 'z' {
				myarray[i] = 'A' - 'a' + myarray[i]
			}
			isString = false
		} else if myarray[i] >= 'A' && myarray[i] <= 'Z' {
			myarray[i] = 'a' - 'A' + myarray[i]
		} else if checker(myarray[i]) {
			isString = true
		}
	}
	return string(myarray)
}

func checker(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}
