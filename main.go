package main

import (
	"fmt"
	"os"
)

func myAtoi(s string) int {
	i := 0
	size := len(s)
	if size == 0 {
		return 0
	}
	isNegative := 1
	result := 0
	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			isNegative = -1
		}
		i++
	}
	for i < size {
		if s[i] >= '0' && s[i] <= '9' {
			tmp := int(s[i] - '0')
			result = result*10 + tmp
		} else {
			return 0
		}
		i++
	}
	return result * isNegative
}

func makeCopy(args []string) [][]int {
	Sudoku := make([][]int, len(args))
	for i := 0; i < len(args); i++ {
		Sudoku[i] = make([]int, len(args[i]))
		for j := 0; j < len(args[i]); j++ {
			Sudoku[i][j] = myAtoi(string(args[i][j]))
		}
	}
	return Sudoku
}

func printSudoku(Sudoku [][]int) {
	fmt.Printf("\n")
	for i := 0; i < len(Sudoku); i++ {
		fmt.Println(Sudoku[i])
	}
	fmt.Printf("\n")
}

//-------------------------------------------------------------------------------------

func canPut(sudoku [][]int, col int, row int, value int) bool {
	return !ifVertical(sudoku, col, row, value) &&
		!ifHorizontal(sudoku, col, row, value) &&
		!ifSquare(sudoku, col, row, value)
}

func ifVertical(sudoku [][]int, col int, row int, value int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[i][col] == value {
			return true
		}
	}
	return false
}

func ifHorizontal(sudoku [][]int, col int, row int, value int) bool {
	for i := range [9]int{} {
		if sudoku[row][i] == value {
			return true
		}
	}
	return false
}

func ifSquare(sudoku [][]int, col int, row int, value int) bool {
	scol, srow := int(col/3)*3, int(row/3)*3
	for drow := range [3]int{} {
		for dcol := range [3]int{} {
			if sudoku[srow+drow][scol+dcol] == value {
				return true
			}
		}
	}
	return false
}

func nextCoordinates(col int, row int) (int, int) {
	nextCol, nextRow := (col+1)%9, row
	if nextCol == 0 {
		nextRow = row + 1
	}
	return nextCol, nextRow
}

func solve(sudoku [][]int, col int, row int) bool {
	if row == 9 {
		return true
	}
	if sudoku[row][col] != 0 {
		return solve(nextCoordinates(col, row))
	} else {
		for i := range [9]int{} {
			v := i + 1
			if canPut(col, row, v) {
				sudoku[row][col] = v
				if solve(nextCoordinates(col, row)) {
					return true
				}
				sudoku[row][col] = 0
			}
		}
		return false
	}
}

func main() {
	args := os.Args[1:]
	Sudoku := makeCopy(args)
	// Sudoku[8][1] = 5
	if len(args) == 9 {
		printSudoku(Sudoku)
	}
}
