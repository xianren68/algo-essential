package __numberOfIslands

func numIslands(grid [][]byte) int {
	res := 0
	for i, b := range grid {
		for j, b2 := range b {
			if b2 == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

// 扩散
func dfs(grid [][]byte, i, j int) {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = 0
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
