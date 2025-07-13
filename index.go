package piscine

func Index(s string, toFind string) int {
	lengthOfS := len(s)
	lengthOfX := len(toFind)
	if lengthOfX == 0 || lengthOfX > lengthOfS {
		return -1
	}
	for index := 0; index <= lengthOfS-lengthOfX; index++ {
		if s[index:index+lengthOfX] == toFind {
			return index
		}
	}
	return -1
}
