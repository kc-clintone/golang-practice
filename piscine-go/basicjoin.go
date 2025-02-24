package piscine

func BasicJoin(elems []string) string {
	res := ""
	for _, chars := range elems {
		res += chars
	}
	return res
}
