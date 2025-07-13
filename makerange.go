package piscine

func MakeRange(min, max int) []int {
	distance := max - min
	if min < max {
		res := make([]int, distance)
		for i := range res {
			res[i] = min + i
		}
		return res
	}
	return nil
}
