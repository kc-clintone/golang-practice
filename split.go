package piscine

func Split(s, sep string) []string {
	result := ""
	count := 0
	var lengthS []string
	t := len(s) - len(sep)
	for i := range s {
		if sep != string(s[t]) && len(s)-len(sep) < i {
			result = result + string(s[i])
		}
		if count <= 0 && len(s)-len(sep) >= i {
			if sep != s[i:len(sep)+i] {
				result = result + string(s[i])
			} else if sep == s[i:len(sep)+i] {
				count = len(sep)
				lengthS = append(lengthS, result)
				result = ""
			}
		}
		count--
	}
	lengthS = append(lengthS, result)
	return lengthS
}
