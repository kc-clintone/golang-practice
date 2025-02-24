package main

import (
	"os"
)

const (
	CONST1 = 9223372036854775807
	CONST2 = 9223372036854775808
)

func resultValid(arg string, a, b int) bool {
	var result uint
	isNegative := false
	if a < 0 {
		a = -a
	} else if b < 0 {
		b = -b
	}
	if arg == "+" {
		result = uint(a) + uint(b)
		if a < 0 && b < 0 {
			isNegative = true
		}
	} else if arg == "-" {
		result = uint(a) + uint(b)
		if a < 0 && b < 0 {
			isNegative = true
		}
	} else if arg == "*" {
		var maxValue int
		var minValue int
		if a > b {
			maxValue = a
			minValue = b
		} else if a < b {
			maxValue = b
			minValue = a
		}
		for i := 0; i < maxValue; i++ {
			result += uint(a) * uint(minValue)
			if result > CONST1 || result > CONST2 && isNegative {
				return false
			}
		}
		if a < 0 || b < 0 {
			isNegative = true
		}
	}
	if result > CONST1 || result > CONST2 && isNegative {
		return false
	}
	return true
}

func operation(arg string, a, b int) string {
	var result int
	var newString string
	var save int
	var signeString string

	if arg == "+" {
		if !resultValid(arg, a, b) {
			return newString
		}
		result = a + b
	} else if arg == "-" {
		if !resultValid(arg, a, b) {
			return newString
		}
		result = a - b
	} else if arg == "*" {
		if !resultValid(arg, a, b) {
			return newString
		}
		result = a * b
	} else if arg == "/" {
		if b == 0 {
			os.Stderr.WriteString("No division by 0\n")
			return newString
		} else {
			result = a / b
		}
	} else if arg == "%" {
		if b == 0 {
			os.Stderr.WriteString("No modulo by 0\n")
			return newString
		} else {
			result = a % b
		}
	}

	if result == 0 {
		return "0"
	}
	if result < 0 {
		signeString = "-"
		result = -result
	}
	for result > 0 {
		save = result % 10
		newString = string(save+48) + newString
		result /= 10
	}
	if signeString != "" {
		newString = signeString + newString
	}
	return newString
}

func convertStringNumber(arg string) int {
	counter := 1
	var num int
	isNegative := false

	for i := len(arg) - 1; i >= 0; i-- {
		if i == 0 && arg[i] == 45 {
			isNegative = true
		} else {
			num += int(arg[i]-48) * counter
			counter *= 10
		}
	}
	if isNegative {
		return -num
	}
	return num
}

func validValue(arg string) bool {
	counter := 1
	isNegative := false
	argSave := arg

	var num uint
	for i := 0; i < len(arg); i++ {
		if arg[0] != 45 && arg[i] < 48 || arg[i] > 57 {
			return false
		}
	}
	if arg[0] == 45 {
		isNegative = true
		argSave = arg[1:]
	}
	for i := len(argSave) - 1; i >= 0; i-- {
		num += uint(argSave[i]-48) * uint(counter)
		counter *= 10
	}
	if num > CONST1 || num > CONST2 && isNegative {
		return false
	}
	return true
}

func validOperator(arg string) bool {
	if len(arg) == 1 {
		if arg == "+" || arg == "-" || arg == "*" || arg == "/" || arg == "%" {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) == 4 {
		os.Args = os.Args[1:]
		for i := 0; i < len(os.Args); i++ {
			if i != 1 {
				if !validValue(os.Args[i]) {
					return
				}
			} else {
				if !validOperator(os.Args[i]) {
					return
				}
			}
		}
		seultsAt := operation(os.Args[1], convertStringNumber(os.Args[0]), convertStringNumber(os.Args[2]))
		os.Stderr.WriteString(seultsAt)
		if seultsAt != "" {
			os.Stderr.WriteString("\n")
		}
	}
}
