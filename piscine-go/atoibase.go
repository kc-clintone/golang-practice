package piscine

func AtoiBase(s string, t string) int {
	result := 0
	lenthS := 0
	myarr := map[rune]bool{}
	for _, i := range t {
		if myarr[i] || i == '-' || i == '+' {
			return result
		}
		myarr[i] = true
		lenthS++
	}
	if lenthS > 1 {
		for _, target := range s {
			count := 0
			if myarr[target] {
				for _, j := range t {
					if j == target {
						break
					}
					count++
				}
				result = result*lenthS + count
			}
		}
	}
	return result
}
