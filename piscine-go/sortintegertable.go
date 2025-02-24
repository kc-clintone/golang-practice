package piscine

func SortIntegerTable(table []int) {
	k := len(table)
	if k < 2 {
		return
	}
	for i := 0; i < k-1; i++ {
		sorted := false
		for j := 0; j < k-1-i; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
				sorted = true
			}
		}
		if !sorted {
			break
		}
	}
}
