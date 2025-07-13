package main

import "github.com/01-edu/z01"

func main() {
	lower := "zyxwvutsrqponmlkjihgfedcba"
	counter := 0
	for _, c := range lower {
		z01.PrintRune(c)
		counter++
	}
	z01.PrintRune('\n')
}
