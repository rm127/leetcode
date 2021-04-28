package main

func isValidSudoku(board [][]byte) bool {
	for row := range board {
		for column := range board[row] {
			current := board[row][column]
			if current != '.' {
				for i := 0; i < 9; i++ {
					if i != column && board[row][i] == current {
						return false
					}
					if i != row && board[i][column] == current {
						return false
					}
					if !(row == row/3*3+i%3 && column == column/3*3+i/3) && board[row/3*3+i%3][column/3*3+i/3] == current {
						return false
					}
				}
			}
		}
	}
	return true
}
