package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	number := 0
	for i, v := range str {
		if v == ' ' {
			summary[str[number:i]] += 1
			number = i + 1
		}
	}
	summary[str[number:]] += 1
	return summary
}
