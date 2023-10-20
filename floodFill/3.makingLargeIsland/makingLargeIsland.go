package __makingLargeIsland

// 记录每个岛屿的大小
func largestIsland(grid [][]int) int {
	// 记录所有岛屿的大小
	islandArea := make(map[int]int)
	// 记录所有岛屿的最大值
	ans := 1
	// 将岛屿的0修改为id
	id := 2
	for i, val := range grid {
		for j, v := range val {
			if v == 1 {
				area := dfs(grid, i, j, id)
				if ans < area {
					ans = area
				}
				islandArea[id] = area
				id++
			}
		}
	}
	var res int
	// 用于去重
	size := make([]bool, id)
	var up, down, left, right int
	// 每个0扩散向四周看能连到几片
	for i, val := range grid {
		for j, v := range val {

			if v == 0 {
				res = 1
				// 上
				if i > 0 && grid[i-1][j] != 0 {
					up = grid[i-1][j]
					if !size[up] {
						res += islandArea[up]
						size[up] = true
					}
				}
				// 下
				if i < len(grid)-1 && grid[i+1][j] != 0 {
					down = grid[i+1][j]
					if !size[down] {
						res += islandArea[down]
						size[down] = true
					}
				}
				// 左
				if j > 0 && grid[i][j-1] != 0 {
					left = grid[i][j-1]
					if !size[left] {
						res += islandArea[left]
						size[left] = true
					}
				}
				// 右
				if j < len(val)-1 && grid[i][j+1] != 0 {
					right = grid[i][j+1]
					if !size[right] {
						res += islandArea[right]
						size[right] = true
					}
				}
				if res > ans {
					ans = res
				}
				size[up] = false
				size[down] = false
				size[left] = false
				size[right] = false

			}
		}
	}
	return ans
}

// 返回岛屿大小
func dfs(grid [][]int, i, j, id int) int {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] != 1 {
		return 0
	}
	res := 1
	grid[i][j] = id
	res += dfs(grid, i-1, j, id)
	res += dfs(grid, i+1, j, id)
	res += dfs(grid, i, j-1, id)
	res += dfs(grid, i, j+1, id)
	return res

}
