package piscine

func TrimAtoi(s string) int {
	trimmed := 0
	sign := 1

	for _, x := range s {
		if x == '-' && trimmed == 0 {
			sign = -1
		} else if x >= '0' && x <= '9' {
			trimmed = trimmed*10 + int(x-'0')
		}
	}
	return trimmed * sign
}
