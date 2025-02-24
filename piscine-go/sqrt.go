package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	for k := 0; k*k <= nb; k++ {
		if k*k == nb {
			return k
		}
	}
	return 0
}
