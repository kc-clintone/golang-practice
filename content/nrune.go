package piscine

func NRune(s string, n int) rune {
	myrunes := []rune(s)
	if n <= 0 || n > len(myrunes) {
		return 0
	}
	return myrunes[n-1]
}
