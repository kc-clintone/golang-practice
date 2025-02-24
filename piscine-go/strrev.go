package piscine

func StrRev(s string) string {
	reversedStr := []rune(s)
	for i, j := 0, len(reversedStr)-1; i < j; i, j = i+1, j-1 {
		reversedStr[i], reversedStr[j] = reversedStr[j], reversedStr[i]
	}
	return string(reversedStr)
}
