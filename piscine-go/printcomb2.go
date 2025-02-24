package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := 1; j <= 99; j++ {
			numbers := []rune{
				'0' + rune(i/10), '0' + rune(i%10), ' ',
				'0' + rune(j/10), '0' + rune(j%10),
			}
			for _, digits := range numbers {
				z01.PrintRune(digits)
			}
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
