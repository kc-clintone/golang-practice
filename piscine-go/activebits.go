package piscine

func ActiveBits(n int) int {
	counter := 0
	for n > 0 {
		if n & 1 == 1 {
			counter++
		}
		n>>=1
	}
	return counter
}
