package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			ascending = false
		}
		if f(a[i-1], a[i]) < 0 {
			descending = false
		}
	}
	return ascending || descending
}
