package piscine

func BasicAtoi(s string) int {
	convertedStr := 0

	for _, c := range s {
		convertedStr = convertedStr*10 + int(c-'0')
	}
	return convertedStr
}
