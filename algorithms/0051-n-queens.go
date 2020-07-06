package algorithms

import "bytes"

func solveNQueens(n int) [][]string {
	solutions := [][]string{}
	board := generateBoard(n)
	flags := generateFlags(n)
	solve(&solutions, board, flags, n, 0)
	return solutions
}

func generateBoard(n int) [][]bool {
	b := make([][]bool, n)
	for i := range b {
		b[i] = make([]bool, n)
	}
	return b
}

func generateFlags(n int) [][]bool {
	return [][]bool{
		make([]bool, n),
		make([]bool, n),
		make([]bool, n*2-1),
		make([]bool, n*2-1),
	}
}

func solve(solutions *[][]string, board [][]bool, flags [][]bool, n int, row int) {
	// fmt.Printf("======== DEBUG: %d ========\n", row)
	// printBoard(board)
	if row == n {
		*solutions = append(*solutions, generateSolution(board))
		return
	}
	for p := row * n; p < (row+1)*n; p++ {
		if !get(flags, p, n) {
			set(board, flags, p, n, true)
			solve(solutions, board, flags, n, row+1)
			set(board, flags, p, n, false)
		}
	}
}

func getValidPoints(flags [][]bool, n int) []int {
	var p []int
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if flags[0][x] || flags[1][y] || flags[2][x+y] || flags[3][x-y+n-1] {
				continue
			}
			p = append(p, x*n+y)
		}
	}
	return p
}

func generateSolution(board [][]bool) []string {
	var solution []string
	for x := range board {
		var buff bytes.Buffer
		for y := range board[x] {
			if board[x][y] {
				buff.WriteByte('Q')
			} else {
				buff.WriteByte('.')
			}
		}
		solution = append(solution, buff.String())
	}
	return solution
}

func get(flags [][]bool, p int, n int) bool {
	x, y := p/n, p%n
	return flags[0][x] || flags[1][y] || flags[2][x+y] || flags[3][x-y+n-1]
}

func set(board [][]bool, flags [][]bool, p int, n int, v bool) {
	x, y := p/n, p%n
	board[x][y] = v
	flags[0][x] = v
	flags[1][y] = v
	flags[2][x+y] = v
	flags[3][x-y+n-1] = v
}
