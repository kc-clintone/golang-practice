package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	myslice := []rune(base)
	var n int
	x := len(myslice)
	var r string
	index := []int{}
	if len(myslice) < 2 {
		printstr("NV")
		return
	}
	for i := 0; i < len(myslice); i++ {
		for j := 0; j < len(myslice); j++ {
			if i != j && myslice[i] == myslice[j] {
				printstr("NV")
				return
			}
		}
		if myslice[i] == '+' || myslice[i] == '-' {
			printstr("NV")
			return
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
		n = -nbr
	} else if nbr > 0 {
		n = nbr
	} else {
		z01.PrintRune('0')
	}
	if nbr == -9223372036854775808 {
		printstr("9223372036854775808")
	}
	for n > 0 {
		index = append(index, n%x)
		n /= x
	}
	for i := 0; i < len(index); i++ {
		r = string(myslice[index[i]]) + r
	}
	printstr(r)
}

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
