package piscine

func returnNbrBase(nbr int, base string) string {
	result := ""
	myrune := []rune(base)
	IsNegative := false
	reset := 0
	if nbr < 0 {
		IsNegative = true
		reset = (nbr % len(myrune)) * -1
		result = string(myrune[reset]) + result
		nbr = (nbr / len(myrune)) * -1
	}
	if nbr < len(myrune) {
		result = string(myrune[nbr]) + result
	} else {
		for nbr > 0 {
			reset = nbr % len(myrune)
			result = string(myrune[reset]) + result
			nbr = (nbr / len(myrune))
		}
	}
	if IsNegative {
		result = "-" + result
	}
	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return returnNbrBase(AtoiBase(nbr, baseFrom), baseTo)
}
