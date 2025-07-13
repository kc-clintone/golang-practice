package piscine

func ToUpper(s string) string {
	res := ""
	for _, x := range s {
		if x >= 'a' && x <= 'z' {
			res += string(x - 32)
		} else {
			res += string(x)
		}
	}
	return res
}
