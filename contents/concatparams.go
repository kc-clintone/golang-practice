package piscine

func ConcatParams(args []string) string {
	result := ""
	for i := range args {
		result = result + args[i]
		if i != len(args)-1 {
			result = result + "\n"
		}
	}
	return result
}
