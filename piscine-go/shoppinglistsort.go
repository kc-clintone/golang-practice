package piscine

func ShoppingListSort(slice []string) []string {
	for i := 1; i < len(slice); i++ {
		counter := slice[i]
		j := i - 1
		for j >= 0 && len(slice[j]) > len(counter) {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = counter
	}
	return slice
}
