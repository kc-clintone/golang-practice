package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0
	res := 1
	if len(os.Args) > 1 {
		mynum := os.Args[1:]
		if mynum[0] == "--upper" {
			mynum = mynum[1:]
			for i := range mynum {
				res = 1
				for range mynum[i] {
					res = res * 10
				}
				count = 0
				for j := range mynum[i] {
					res = res / 10
					count = count + (int(mynum[i][j])-48)*res
				}
				z01.PrintRune(rune(count + 64))
			}
		} else if len(mynum) > 0 {
			for I := range mynum {
				res = 1
				for range mynum[I] {
					res = res * 10
				}
				count = 0
				for J := range mynum[I] {
					res = res / 10
					count = count + (int(mynum[I][J])-48)*res
				}
				if count >= 27 {
					z01.PrintRune(' ')
				} else {
					z01.PrintRune(rune(count + 96))
				}
			}
		}
		z01.PrintRune(10)
	}
}
