package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	isNegative := 1
	counter := 0
	convertedStr := 0

	if s[0] == '-' {
		isNegative = -1
		counter = 1
	} else if s[0] == '+' {
		counter = 1
	}

	for i := counter; i < len(s); i++ {
		char := s[i]
		if char < '0' || char > '9' {
			return 0
		}
		convertedStr = convertedStr*10 + int(char-'0')
	}
	return convertedStr * isNegative
}
