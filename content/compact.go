package piscine

func Compact(ptr *[]string) int{
	compacted := []string{}
	for _, char := range *ptr {
		if char != "" {
			compacted = append(compacted, char)
		}
	}
	*ptr = compacted
	return len(compacted)
}