package __isLands

/* leetcode 200.岛屿数量 */

func numIslands(grid [][]byte) int {
	u := build(len(grid) * len(grid[0]))
	l := len(grid[0])
	for i, val := range grid {
		for j, v := range val {
			if v == '0' {
				u.set--
				continue
			}
			// 寻找左边的1
			if j > 0 {
				if grid[i][j-1] == '1' {
					// 合并集合
					u.union(index(i, j, l), index(i, j-1, l))
				}
			}
			// 寻找上边的1
			if i > 0 {
				if grid[i-1][j] == '1' && !(u.isSameset(index(i, j, l), index(i-1, j, l))) {
					// 合并集合
					u.union(index(i, j, l), index(i-1, j, l))
				}
			}
		}
	}
	return u.set
}

// 计算对应的索引
func index(i, j, l int) int {
	return i*l + j
}

type unionFind struct {
	father []int
	set    int // 集合的总数量
}

func build(n int) *unionFind {
	u := &unionFind{
		father: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.father[i] = i
		u.set++
	}
	return u
}
func (u *unionFind) find(n int) int {
	if u.father[n] == n {
		return n
	}
	u.father[n] = u.find(u.father[n])
	return u.father[n]
}
func (u *unionFind) isSameset(x, y int) bool {
	return u.find(x) == u.find(y)
}
func (u *unionFind) union(x, y int) {
	f1, f2 := u.find(x), u.find(y)
	u.father[f1] = f2
	u.set--
}
