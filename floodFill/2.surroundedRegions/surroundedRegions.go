package __surroundedRegions

// 从四个边界的'o'开始扩散，那些不能被扩散到的即为被包围的
func solve(board [][]byte) {
	col := len(board[0])
	row := len(board)
	// 上下边界
	for i := 0; i < col; i++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
		if board[row-1][i] == 'O' {
			dfs(board, row-1, i)
		}
	}
	// 左右边界
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][col-1] == 'O' {
			dfs(board, i, col-1)
		}

	}
	// 恢复并将被包围的改成'X'
	for i, val := range board {
		for j, v := range val {
			if v == 0 {
				// 恢复边界
				board[i][j] = 'O'
				continue
			}
			if v == 'O' {
				// 修改包围
				board[i][j] = 'X'
			}

		}
	}
}
func dfs(board [][]byte, i, j int) {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 0
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
}
