package main

import (
	"fmt"
	"os"
)

func main() {
	len := len(os.Args)
	if len < 2 {
		fmt.Println("File name missing")
		return
	}
	if len > 2 {
		fmt.Println("Too many arguments")
		return
	}
	file := os.Args[1]
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(string(data))
}
