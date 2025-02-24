package piscine

func UltimateDivMod(a *int, b *int) {
	if *b == 0 || b == nil {
		return
	}
	quotient := *a / *b
	modulo := *a % *b

	*a = quotient
	*b = modulo
}
