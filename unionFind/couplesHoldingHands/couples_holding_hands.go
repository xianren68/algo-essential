package couplesholdinghands

/* leetcode 765.情侣牵手 */
func minSwapsCouples(row []int) int {
	u := build(len(row) / 2)
	for i := 0; i < len(row); i += 2 {
		r1, r2 := row[i]/2, row[i+1]/2
		// 当前两个位置不是情侣
		if r1 != r2 {
			if !u.isSameset(r1, r2) {

				// 合并两个集合
				u.union(r1, r2)
			}

		}
	}
	return len(row)/2 - u.set
}

type unionFind struct {
	father []int
	set    int // 集合的总数量
}

func build(n int) *unionFind {
	u := &unionFind{
		father: make([]int, n),
	}
	u.set = n
	for i := 0; i < n; i++ {
		u.father[i] = i
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
