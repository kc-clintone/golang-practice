package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		for _, file := range os.Args[1:] {
			data, err := os.ReadFile(file)
			if err != nil {
				printErr(file)
				os.Exit(1)
			}
			printFile(string(data))
		}
	} else {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return
		}
		printFile(string(data))
	}
}

func printFile(data string) {
	for _, char := range data {
		z01.PrintRune(char)
	}
}

func printErr(file string) {
	errMsg := "ERROR: open " + file + ": no such file or directory\n"
	for _, char := range errMsg {
		z01.PrintRune(char)
	}
}
