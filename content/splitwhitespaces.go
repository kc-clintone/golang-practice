package piscine

func SplitWhiteSpaces(s string) []string {
	splitted := ""
	isSpace := true
	var length []string
	for _, i := range s {
		if !(rune(i) == ' ' || rune(i) == '\n' || rune(i) == '	') {
			splitted = splitted + string(i)
			isSpace = false
		} else {
			if isSpace == false && splitted != "" {
				length = append(length, splitted)
				splitted = ""
			} else {
				isSpace = true
			}
		}
	}
	if isSpace == false && splitted != "" {
		length = append(length, splitted)
	}
	return length
}
