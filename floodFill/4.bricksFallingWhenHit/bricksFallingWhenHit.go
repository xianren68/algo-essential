package __bricksFallingWhenHit

// 矩阵行列数
var r, c int

// 矩阵
var g [][]int

func hitBricks(grid [][]int, hits [][]int) []int {
	r, c = len(grid)-1, len(grid[0])-1
	g = grid
	res := make([]int, len(hits))
	// 所有砖块都在天花板上，不会有砖块掉落
	if r == 1 {
		return res
	}
	// 去掉所有打中的位置
	for _, hit := range hits {
		grid[hit[0]][hit[1]]--
	}
	// 将所有与天花板连接的1都改成2
	for i := 0; i <= c; i++ {
		dfs(0, i)
	}
	// 时光倒流，倒序恢复打掉的砖块，判断其上下左右是否有2,若有，则感染1，看能将几个1变成2
	for i := len(hits) - 1; i >= 0; i-- {
		j, z := hits[i][0], hits[i][1]
		g[j][z]++
		if worth(j, z) {
			res[i] = dfs(j, z) - 1
		}
	}
	return res
}

// 感染，将1变成2
func dfs(i, j int) int {
	if i < 0 || i == r || j < 0 || j == c || g[i][j] != 1 {
		return 0
	}
	g[i][j] = 2
	return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
}

// 判断是否有必要感染
// 炮弹不为空 && (打在天花板 || 上下左右有2
func worth(i, j int) bool {
	return g[i][j] == 1 &&
		(i == 0 ||
			(i > 0 && g[i-1][j] == 2) ||
			(i < c && g[i+1][j] == 2) ||
			(j > 0 && g[i][j-1] == 2) ||
			(j < r && g[i][j+1] == 2))

}
