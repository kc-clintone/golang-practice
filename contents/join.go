package piscine

func Join(strs []string, sep string) string {
	results := ""

	for i := 0; i < len(strs); i++ {
		results += strs[i]
		if i < len(strs)-1 {
			results += sep
		}
	}
	return results
}
