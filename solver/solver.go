package solver

func CanIPut(sudoku [][]int, col int, row int, value int) bool {
	return !CheckVertical(sudoku, col, row, value) && !CheckHorizontal(sudoku, col, row, value) && !CheckSquare3_3(sudoku, col, row, value)
}

func CheckVertical(sudoku [][]int, col int, row int, value int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[i][col] == value {
			return true
		}
	}
	return false
}

func CheckHorizontal(sudoku [][]int, col int, row int, value int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[row][i] == value {
			return true
		}
	}
	return false
}

func CheckSquare3_3(sudoku [][]int, col int, row int, value int) bool {
	scol, srow := int(col/3)*3, int(row/3)*3
	for drow := 0; drow < 3; drow++ {
		for dcol := 0; dcol < 3; dcol++ {
			if sudoku[srow+drow][scol+dcol] == value {
				return true
			}
		}
	}
	return false
}

func NextEmptyCell(col int, row int) (int, int) {
	nextCol, nextRow := (col+1)%9, row
	if nextCol == 0 {
		nextRow = row + 1
	}
	return nextCol, nextRow
}

func SudokuSolver(sudoku [][]int, col int, row int) bool {
	if row == 9 {
		return true
	}
	if sudoku[row][col] != 0 {
		tmpx, tmpy := NextEmptyCell(col, row)
		return SudokuSolver(sudoku, tmpx, tmpy)
	} else {
		for i := 0; i < 9; i++ {
			value := i + 1
			if CanIPut(sudoku, col, row, value) {
				sudoku[row][col] = value
				tmpx, tmpy := NextEmptyCell(col, row)
				if SudokuSolver(sudoku, tmpx, tmpy) {
					return true
				}
				sudoku[row][col] = 0
			}
		}
		return false
	}
}
