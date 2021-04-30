package main

func solveSudoku(board [][]byte) {
	numbers := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var options [9][9][]byte

	// fill with possible numbers
	for row := range options {
		for column := range options[row] {
			options[row][column] = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
		}
	}

	// update options for all numbers in input
	for row := range board {
		for column := range board[row] {
			if board[row][column] != '.' {
				updateOptions(row, column, board, &options)
			}
		}
	}

	prevChange, currentChange := 0, 1
	for prevChange < currentChange {
		prevChange = currentChange

		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				// for each empty cell on board
				if board[r][c] == '.' {

					// loop through all options for the cell
					for o := 0; o < len(options[r][c]); o++ {
						aloneX, aloneY, aloneB := true, true, true

						// loop over all cells in row/column/box
						for i := 0; i < len(numbers); i++ {

							// check if the options for this cell contains
							// the given option from the starting cell
							// if this is the case, then we know it isn't
							// the only instance of this option in the row/column/box

							// check row
							if board[r][i] == '.' && c != i && arrayContains(&options[r][i], options[r][c][o]) {
								aloneX = false
							}
							// check column
							if board[i][c] == '.' && r != i && arrayContains(&options[i][c], options[r][c][o]) {
								aloneY = false
							}

							// check box
							boxRow := r/3*3 + i%3
							boxColumn := c/3*3 + i/3
							if board[boxRow][boxColumn] == '.' && !(r == boxRow && c == boxColumn) &&
								arrayContains(&options[boxRow][boxColumn], options[r][c][o]) {
								aloneB = false
							}
						}

						// if there exists a row/column/box where the option is alone,
						// then r,c is the only possible location for said option
						if aloneX || aloneY || aloneB {
							board[r][c] = options[r][c][o]
							updateOptions(r, c, board, &options)
							currentChange++
							break
						}
					}
				}
			}
		}
	}
}

func updateOptions(row int, column int, board [][]byte, options *[9][9][]byte) {
	for i := 0; i < 9; i++ {
		removeFromArray(&options[row][i], board[row][column])
		removeFromArray(&options[i][column], board[row][column])
		removeFromArray(&options[row/3*3+i%3][column/3*3+i/3], board[row][column])
	}
	// if only one option available then insert it
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if len(options[r][c]) == 1 && board[r][c] == '.' {
				board[r][c] = options[r][c][0]
				// and update
				updateOptions(r, c, board, options)
			}
		}
	}
}

func arrayContains(s *[]byte, match byte) bool {
	for _, a := range *s {
		if a == match {
			return true
		}
	}
	return false
}

func removeFromArray(s *[]byte, match byte) {
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == match {
			(*s)[i] = (*s)[len(*s)-1]
			*s = (*s)[:len(*s)-1]
		}
	}
}
