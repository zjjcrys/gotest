package leedcode

//130 先把边界以及和边界相连的剔除，
func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if (i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1) && board[i][j] == 'O' {
				solveDFS(&board, i, j)
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '-' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func solveDFS(board *[][]byte, row int, col int) {
	(*board)[row][col] = '-'
	if row-1 >= 0 && (*board)[row-1][col] == 'O' {
		solveDFS(board, row-1, col)
	}
	if row+1 < len(*board) && (*board)[row+1][col] == 'O' {
		solveDFS(board, row+1, col)
	}
	if col-1 >= 0 && (*board)[row][col-1] == 'O' {
		solveDFS(board, row, col-1)
	}
	if col+1 < len((*board)[0]) && (*board)[row][col+1] == 'O' {
		solveDFS(board, row, col+1)
	}
}
