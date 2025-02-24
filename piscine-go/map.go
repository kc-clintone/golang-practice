package piscine

func Map(f func(int) bool, a []int) []bool {
	mySlice := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			mySlice[i] = true
		} else {
			mySlice[i] = false
		}
	}
	return mySlice
}
