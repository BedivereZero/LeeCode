// package algorithms
package main

import (
	"errors"
	"fmt"
)

func main() {
	b := [][]byte{{'.', '8', '7', '6', '5', '4', '3', '2', '1'}, {'2', '.', '.', '.', '.', '.', '.', '.', '.'}, {'3', '.', '.', '.', '.', '.', '.', '.', '.'}, {'4', '.', '.', '.', '.', '.', '.', '.', '.'}, {'5', '.', '.', '.', '.', '.', '.', '.', '.'}, {'6', '.', '.', '.', '.', '.', '.', '.', '.'}, {'7', '.', '.', '.', '.', '.', '.', '.', '.'}, {'8', '.', '.', '.', '.', '.', '.', '.', '.'}, {'9', '.', '.', '.', '.', '.', '.', '.', '.'}}
	fmt.Println(isValidSudoku(b))
}

func isValidSudoku(board [][]byte) bool {
	return valid(board, 0, 0)
}

func next(board [][]byte, x int, y int) (int, int, error) {
	for i := x; i < 9; i++ {
		for j := y; j < 9; j++ {
			if board[i][j] == '.' {
				return i, j, nil
			}
		}
	}
	return 0, 0, errors.New("not found")
}

func valid(board [][]byte, x int, y int) bool {
	for i := range board {
		for j := range board[i] {
			fmt.Printf(" %c", board[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("valid: %d, %d\n", x, y)

	x, y, err := next(board, x, y)
	fmt.Printf("next: %d, %d\n", x, y)

	if err != nil {
		return true
	}

	candidate := getCandidate(board, x, y)

	for _, c := range candidate {
		board[x][y] = c
		if valid(board, (x*9+y+1)/9, (x*9+y+1)%9) {
			return true
		}
	}
	board[x][y] = '.'

	return false
}

func getCandidate(board [][]byte, x int, y int) []byte {
	// fmt.Printf("A: x: %d, y: %d\n", x, y)
	flag := make([]bool, 9)
	// same column
	for i := 0; i < 9; i++ {
		if board[i][y] != '.' {
			flag[int(board[i][y]-'0'-1)] = true
		}
	}
	// same row
	for i := 0; i < 9; i++ {
		if board[x][i] != '.' {
			flag[int(board[x][i]-'0'-1)] = true
		}
	}
	// same block
	for i := x / 3 * 3; i < x/3*3+3; i++ {
		for j := y / 3 * 3; j < y/3*3+3; j++ {
			if board[i][j] != '.' {
				// fmt.Printf("B: x: %d, y: %d, f: %d\n", i, j, int(board[i][j]-'0'))
				flag[int(board[i][j]-'0'-1)] = true
			}
		}
	}

	r := []byte{}
	for i, f := range flag {
		if !f {
			r = append(r, byte('0'+i+1))
		}
	}
	return r
}
