package piscine

func IsUpper(s string) bool {
	for _, character := range s {
		if character < 'A' || character > 'Z' {
			return false
		}
	}
	return true
}
