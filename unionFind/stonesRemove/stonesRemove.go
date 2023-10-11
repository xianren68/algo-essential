package stonesRemove

/* leetcode 947.移除最多的同行或同列石头 */

func removeStones(stones [][]int) int {
	l := len(stones)
	u := build(l)
	// 记录行列信息
	row := make(map[int]int)
	column := make(map[int]int)
	for i, val := range stones {
		// 判断当前石块行是否有石头
		if _, ok := row[val[0]]; ok {
			// 合并
			u.union(i, row[val[0]])
		} else {
			// 记录到hash表中
			row[val[0]] = i
		}
		// 判断当前列是否有石块
		if _, ok := column[val[1]]; ok {
			u.union(i, column[val[1]])
		} else {
			column[val[1]] = i
		}

	}
	return l - u.sets

}

type unionFind struct {
	father []int
	sets   int
}

func build(n int) *unionFind {
	res := &unionFind{
		father: make([]int, n),
	}
	for i := 0; i < n; i++ {
		res.father[i] = i
	}
	res.sets = n
	return res
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
	if f1 != f2 {
		u.father[f1] = f2
		u.sets--
	}
}
