package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	counter := 0
	myArr := []string{"01", "galaxy", "galaxy 01"}

	for i := range args {
		for _, char := range myArr {
			if args[i] == char {
				counter++
			}
		}
	}
	if counter >= 1 {
		fmt.Println("Alert!!!")
	}
}