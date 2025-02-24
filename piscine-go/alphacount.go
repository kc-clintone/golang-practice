package piscine

func AlphaCount(s string) int {
	result := 0
	for _, characters := range s {
		if (characters >= 'a' && characters <= 'z') || (characters >= 'A' && characters <= 'Z') {
			result++
		}
	}
	return result
}
