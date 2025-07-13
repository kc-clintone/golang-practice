package piscine

func BasicAtoi2(s string) int {
	convertedStr := 0

	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		convertedStr = convertedStr*10 + int(c-'0')
	}
	return convertedStr
}
