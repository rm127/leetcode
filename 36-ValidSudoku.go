package main

func isValidSudoku(board [][]byte) bool {
	err := false

	// box increaser
	boxVerOffset := 0 // goes 0..3..6..0 (columns)
	boxVerMove := 0   // goes 0..3..6 (rows)

	// row/column check
	for i := 0; i < 9; i++ {
		if err {
			break
		}

		boxHorOffset := 0 // goes 0..1..2..3

		// store current numbers in row/column
		var horNums []byte
		var verNums []byte
		var boxNums []byte

		// loop through the row/column
		for j := 0; j < 9; j++ {

			// horizontally
			horItem := board[i][j]

			// check if current item already is present in row
			for n := range horNums {
				if horNums[n] == horItem {
					err = true
					break
				}
			}
			if err {
				break
			}
			// if not empty and not already present, add to memory
			if horItem != '.' {
				horNums = append(horNums, horItem)
			}

			// vertically
			verItem := board[j][i]

			// check if current item already is present in column
			for n := range verNums {
				if verNums[n] == verItem {
					err = true
					break
				}
			}
			if err {
				break
			}
			// if not empty and not already present, add to memory
			if verItem != '.' {
				verNums = append(verNums, verItem)
			}

			// box
			// create x and y coordinates for each box (1..2..3..0..)
			x := j%3 + boxVerOffset
			y := boxHorOffset + boxVerMove

			boxItem := board[x][y]

			// check if current item already is present in box
			for n := range boxNums {
				if boxNums[n] == boxItem {
					err = true
					break
				}
			}
			if err {
				break
			}
			// if not empty and not already present, add to memory
			if boxItem != '.' {
				boxNums = append(boxNums, boxItem)
			}

			// increase the row count as we go
			if j == 2 || j == 5 {
				boxHorOffset++
			}
		}

		// increase and reset the column count
		if i == 2 || i == 5 {
			boxVerOffset = 0
			boxVerMove += 3
		} else {
			boxVerOffset += 3
		}
	}

	return !err
}
