package main

// 回合范围在 0 ~ 怪兽血量之间
func cutOrPoison(cuts, Poison []int, hp int) int {
	res := hp
	l, r := 0, hp
	var m int
	for l <= r {
		m = l + (r-l)>>1
		if f(cuts, Poison, m, hp) {
			res = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return res
}

// 判断这些回合是否能够杀死怪兽
func f(cuts, poison []int, m, hp int) bool {
	// 每一回合判断从当前到最后回合毒杀和刀砍谁的伤害高
	for i, cut := range cuts {
		if i > m {
			return false
		}
		hp -= maxInt(cut, poison[i]*(m-i))
		if hp <= 0 {
			return true
		}
	}
	return false

}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
