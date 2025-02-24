package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	counter := 1
	for nb > 0 {
		counter *= nb
		nb--
	}
	return counter
}
