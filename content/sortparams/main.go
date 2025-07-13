package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myargs := os.Args[1:]
	for i := 0; i < len(myargs); i++ {
		for j := i + 1; j < len(myargs); j++ {
			if myargs[i] > myargs[j] {
				myargs[i], myargs[j] = myargs[j], myargs[i]
			}
		}
	}
	for _, arg := range myargs {
		for _, k := range arg {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}
