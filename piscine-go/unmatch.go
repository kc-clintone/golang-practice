package piscine

func Unmatch(a []int) int {
	for i, j := range a {
		counter := 1
		for k, l := range a {
			if j == l && k != i {
				counter++
			}
		}
		if counter%2 == 1 {
			return j
		}
	}
	return -1
}
