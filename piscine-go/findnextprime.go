package piscine

func FindNextPrime(nb int) int {
	IsPrime := func(nb int) bool {
		if nb < 2 {
			return false
		}
		if nb == 2 || nb == 3 {
			return true
		}
		if nb%3 == 0 || nb%2 == 0 {
			return false
		}
		for a := 5; a*a <= nb; a += 6 {
			if nb%(a+2) == 0 || nb%a == 0 {
				return false
			}
		}
		return true
	}
	if nb < 2 {
		return 2
	}
	if nb%2 == 0 && nb > 2 {
		nb++
	}
	for !IsPrime(nb) {
		nb += 2
	}
	return nb
}
