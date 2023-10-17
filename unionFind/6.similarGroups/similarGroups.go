package __similarGroups

/* leetcode 117.相似字符串组 */

func numSimilarGroups(strs []string) int {
	u := build(len(strs))
	for i := 0; i < len(strs)-1; i++ {
		for j := i + 1; j < len(strs); j++ {
			// 判断是否在同一组
			if u.isSameset(i, j) {
				continue
			}
			// 判断是否相似
			if isSimilar(strs[i], strs[j]) {
				//合并所在组
				u.union(i, j)
			}
		}
	}
	return u.set
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

// 两个字符串是否相似
func isSimilar(x, y string) bool {
	diff := 0
	for i := 0; i < len(x); i++ {
		if y[i] != x[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	if diff == 0 || diff == 2 {
		return true
	}
	return false
}
