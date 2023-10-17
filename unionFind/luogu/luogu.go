package luogu

type unionFind struct {
	father []int
}

func build(n int) *unionFind {
	u := &unionFind{
		father: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.father[i] = i
	}
	return u
}
func (u *unionFind) find(n int) int {
	if u.father[n] == n {
		return n
	}
	// 递归完成路径压缩
	u.father[n] = u.find(u.father[n])
	return u.father[n]
}
func (u *unionFind) isSameset(x, y int) bool {
	return u.find(x) == u.find(y)
}
func (u *unionFind) union(x, y int) {
	f1, f2 := u.find(x), u.find(y)
	u.father[f1] = f2
}
