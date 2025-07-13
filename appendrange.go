package piscine

func AppendRange(min, max int) []int {
	var size []int
	for i := min; i < max; i++ {
		size = append(size, i)
	}
	return size[0:]
}
