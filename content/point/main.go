package main

import (
	"github.com/01-edu/z01"
)

// func setPoint(ptrx, ptry *int) {
// 	*ptrx = 42
// 	*ptry = 21
// }

// func orintNumber(nb int) {
// 	var results int
// 	var myTable []int
// 	for nb > 0 {
// 		results = nb % 10
// 		myTable = append(myTable, results)
// 		nb /= 10
// 	}
// 	for i := len(myTable) - 1; i >= 0; i-- {
// 		z01.PrintRune(rune(int(myTable[i]) + 48))
// 	}
// }

// func main() {
// 	var x int
// 	var y int
// 	setPoint(&x, &y)
// 	z01.PrintRune('x')
// 	z01.PrintRune(' ')
// 	z01.PrintRune('=')
// 	z01.PrintRune(' ')
// 	orintNumber(x)
// 	z01.PrintRune(',')
// 	z01.PrintRune(' ')
// 	z01.PrintRune('y')
// 	z01.PrintRune(' ')
// 	z01.PrintRune('=')
// 	z01.PrintRune(' ')
// 	orintNumber(y)
// 	z01.PrintRune('\n')
// }
var myArr = "x = 42, y = 21"

func main() {
	for _, i := range myArr {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
