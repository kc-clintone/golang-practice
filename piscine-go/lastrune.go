package piscine

func LastRune(s string) rune {
	myrunes := []rune(s)
	return myrunes[len(myrunes)-1]
}
