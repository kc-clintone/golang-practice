package piscine

func ToLower(s string) string {
	res := ""
	for _, x := range s {
		if x >= 'A' && x <= 'Z' {
			res += string(x + 32)
		} else {
			res += string(x)
		}
	}
	return res
}
