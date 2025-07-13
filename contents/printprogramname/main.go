package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	for i := len(args) - 1; i >= 0; i-- {
		if args[i] == '/' {
			args = args[i+1:]
			break
		}
	}
	for _, k := range args {
		z01.PrintRune(k)
	}
	z01.PrintRune('\n')
}
