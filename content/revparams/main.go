package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myargs := os.Args[1:]
	counter := 0
	for k := range myargs {
		counter = k + 1
	}
	for k := counter - 1; k >= 0; k-- {
		for _, chars := range myargs[k] {
			z01.PrintRune(chars)
		}
		z01.PrintRune('\n')
	}
}
