package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
	}

	res := []int{}
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}

	for x := 0; x < len(res); x++ {
		for y := x + 1; y < len(res); y++ {
			if res[x] > res[y] {
				res[x], res[y] = res[y], res[x]
			}
		}
	}

	for _, digit := range res {
		t := 1
		k := (digit % 10 * t) + '0'
		z01.PrintRune(rune(k))
	}
}
