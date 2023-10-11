package secret

import "sort"

/* leetcode 2092.知晓秘密的专家 */

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	u := build(n)
	// 将0号专家和firstPerson合并
	u.union(0, firstPerson)
	// 将所有会议按时间排序
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	// 循环
	for i := 0; i < len(meetings); {
		// 当前时间
		runtime := meetings[i][2]
		// 遍历当前时间的所有会议
		var j = i
		for ; j < len(meetings) && meetings[j][2] == runtime; j++ {
			x, y := meetings[j][0], meetings[j][1]
			// 合并
			u.union(x, y)
		}
		for t := i; t < j; t++ {
			u.unravel(meetings[t][0])
			u.unravel(meetings[t][1])
		}
		// 跳过这些开过的会
		i = j
	}
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if u.isKnow[u.find(i)] {
			res = append(res, i)
		}
	}
	return res
}

type unionFind struct {
	father []int
	isKnow []bool
}

func build(n int) *unionFind {
	res := &unionFind{
		father: make([]int, n),
		isKnow: make([]bool, n),
	}
	// 设置每个节点为一个集合，代表节点为自己
	for i := 0; i < n; i++ {
		res.father[i] = i
	}
	// 第一位专家知道秘密
	res.isKnow[0] = true
	return res
}

// 查找代表节点
func (u *unionFind) find(n int) int {
	if u.father[n] == n {
		return n
	}
	// 递归查找并压缩路径
	u.father[n] = u.find(u.father[n])
	return u.father[n]
}

// 判断两个节点是否在同一集合中
func (u *unionFind) isSameSet(x, y int) bool {
	return u.father[x] == u.father[y]
}

// 合并集合
func (u *unionFind) union(x, y int) {
	f1, f2 := u.find(x), u.find(y)
	u.father[f2] = f1
	u.isKnow[f1] = u.isKnow[f1] || u.isKnow[f2]

}

// 拆散集合
func (u *unionFind) unravel(n int) {
	// 判断是否在知道秘密的集合内
	// 1.找到代表节点
	f := u.find(n)
	// 2. 判断是否知道秘密
	if !u.isKnow[f] {
		// 将当前节点设置为单节点集合
		u.father[n] = n
	}

}
