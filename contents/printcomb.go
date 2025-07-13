package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for d := '0'; d <= '7'; d++ {
		for b := d + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(d)
				z01.PrintRune(b)
				z01.PrintRune(c)

				if d != '7' || b != '8' || c != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
