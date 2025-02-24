package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	row := 0
	board := [8]int{}
	for k := range board {
		board[k] = -1
	}
	HelperFunc(row, &board)
}

func HelperFunc(row int, board *[8]int) {
	if row == len(board) {
		PrintBits(board)
		// PrintBoard(board)
		return
	}
	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row] = col
			HelperFunc(row+1, board)
			board[row] = -1
		}
	}
}

func isValid(board *[8]int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || board[i]+i == col+row || board[i]-i == col-row {
			return false
		}
	}
	return true
}

func PrintBits(board *[8]int) {
	for _, col := range board {
		z01.PrintRune(rune(col + 1 + '0'))
	}
	z01.PrintRune('\n')
}

// func PrintBoard(board *[8]int) {
// 	for row := 0; row < len(board); row++ {
// 		for col := 0; col < 8; col++ {
// 			if board[row] == col {
// 				z01.PrintRune('Q')
// 			} else {
// 				z01.PrintRune('*')
// 			}
// 		}
// 		z01.PrintRune('\n')
// 	}
// }
