package piscine

func DescendAppendRange(max, min int) []int {
	size := []int{}
	if min >= max {
		return size
	}
	for i := max; i > min; i-- {
		size = append(size, i)
	}
	return size
}
