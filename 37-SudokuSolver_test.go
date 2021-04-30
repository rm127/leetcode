package main

import (
	"bytes"
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{"Example 1", args{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}}, [][]byte{
			{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
			{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
			{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
			{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
			{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
			{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
			{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
			{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
			{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
		}},
		{"Hard 1 from magazine no. 44", args{[][]byte{
			{'2', '.', '.', '.', '8', '1', '7', '.', '.'},
			{'.', '.', '.', '.', '2', '4', '.', '6', '9'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '5'},
			{'6', '.', '7', '.', '.', '.', '.', '.', '.'},
			{'4', '.', '.', '.', '9', '.', '.', '.', '8'},
			{'.', '.', '.', '.', '.', '.', '3', '.', '7'},
			{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'5', '1', '.', '7', '4', '.', '.', '.', '.'},
			{'.', '.', '2', '9', '3', '.', '.', '.', '4'},
		}}, [][]byte{
			{'2', '9', '5', '6', '8', '1', '7', '4', '3'},
			{'3', '7', '8', '5', '2', '4', '1', '6', '9'},
			{'1', '4', '6', '3', '7', '9', '2', '8', '5'},
			{'6', '2', '7', '8', '5', '3', '4', '9', '1'},
			{'4', '5', '3', '1', '9', '7', '6', '2', '8'},
			{'9', '8', '1', '4', '6', '2', '3', '5', '7'},
			{'8', '3', '4', '2', '1', '5', '9', '7', '6'},
			{'5', '1', '9', '7', '4', '6', '8', '3', '2'},
			{'7', '6', '2', '9', '3', '8', '5', '1', '4'},
		}},
		{"Hard 2 from magazine no. 46", args{[][]byte{
			{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '5', '.', '.', '9', '6', '.', '.', '.'},
			{'1', '.', '.', '2', '.', '4', '.', '6', '.'},
			{'.', '3', '.', '9', '.', '.', '2', '1', '7'},
			{'.', '8', '.', '.', '.', '.', '.', '4', '.'},
			{'7', '.', '5', '.', '.', '2', '.', '8', '.'},
			{'.', '6', '.', '5', '.', '7', '.', '.', '8'},
			{'.', '.', '.', '1', '6', '.', '.', '7', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '4'},
		}}, [][]byte{
			{'9', '4', '6', '8', '3', '1', '7', '5', '2'},
			{'8', '5', '2', '7', '9', '6', '4', '3', '1'},
			{'1', '7', '3', '2', '5', '4', '8', '6', '9'},
			{'6', '3', '4', '9', '8', '5', '2', '1', '7'},
			{'2', '8', '1', '6', '7', '3', '9', '4', '5'},
			{'7', '9', '5', '4', '1', '2', '3', '8', '6'},
			{'3', '6', '9', '5', '4', '7', '1', '2', '8'},
			{'4', '2', '8', '1', '6', '9', '5', '7', '3'},
			{'5', '1', '7', '3', '2', '8', '6', '9', '4'},
		}},
		{"First fail from Leetcode", args{[][]byte{
			{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
			{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
			{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
			{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
			{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
			{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
			{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
		}}, [][]byte{
			{'5', '1', '9', '7', '4', '8', '6', '3', '2'},
			{'7', '8', '3', '6', '5', '2', '4', '1', '9'},
			{'4', '2', '6', '1', '3', '9', '8', '7', '5'},
			{'3', '5', '7', '9', '8', '6', '2', '4', '1'},
			{'2', '6', '4', '3', '1', '7', '5', '9', '8'},
			{'1', '9', '8', '5', '2', '4', '3', '6', '7'},
			{'9', '7', '5', '8', '6', '3', '1', '2', '4'},
			{'8', '3', '2', '4', '9', '1', '7', '5', '6'},
			{'6', '4', '1', '2', '7', '5', '9', '8', '3'},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.args.board)

			missingNumbers := 0
			for i := 0; i < len(tt.args.board); i++ {
				for j := 0; j < len(tt.args.board[i]); j++ {
					if tt.args.board[i][j] == '.' {
						missingNumbers++
					}
				}
			}

			for i := 0; i < len(tt.args.board); i++ {
				if bytes.Compare(tt.args.board[i], tt.want[i]) != 0 {
					//t.Errorf("\nFailed at row %v\n%v = isValidSudoku()\n%v = want\nMissing %v numbers", i, string(tt.args.board[i]), string(tt.want[i]), missingNumbers)
					t.Errorf("\n%v = %v\n%v = %v\n%v = %v\n%v = %v\n%v = %v\n%v = %v\n%v = %v\n%v = %v\n%v = %v\nMissing %v numbers",
						string(tt.args.board[0]), string(tt.want[0]),
						string(tt.args.board[1]), string(tt.want[1]),
						string(tt.args.board[2]), string(tt.want[2]),
						string(tt.args.board[3]), string(tt.want[3]),
						string(tt.args.board[4]), string(tt.want[4]),
						string(tt.args.board[5]), string(tt.want[5]),
						string(tt.args.board[6]), string(tt.want[6]),
						string(tt.args.board[7]), string(tt.want[7]),
						string(tt.args.board[8]), string(tt.want[8]),
						missingNumbers)
					break
				}
			}
		})
	}
}
