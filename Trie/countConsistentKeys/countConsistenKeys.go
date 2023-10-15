package countConsistenKeys

import "strconv"

var trie = [150001][11]int{}
var cnt = 1
var pass = [150001]int{}

// Insert 插入前缀树
func Insert(bytes string) {
	cur := 1
	var path byte
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '-' {
			path = 10
		} else if bytes[i] == '#' {
			path = 11
		} else {
			path = bytes[i] - '0'
		}
		pass[cur]++
		// 还不存在此路径，新建
		if trie[cur][path] == 0 {
			trie[cur][path] = cnt + 1
			cnt++
		}
		cur = trie[cur][path]
	}
	pass[cur]++
}

// 查看匹配个数（前缀出现的次数）
func countConsistent(bytes string) int {
	cur := 1
	var path byte
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '-' {
			path = 10
		} else if bytes[i] == '#' {
			path = 11
		} else {
			path = bytes[i] - '0'
		}
		// 还不存在此路径
		if trie[cur][path] == 0 {
			return 0
		}
		cur = trie[cur][path]
	}
	return pass[cur]
}

func countConsistentKeys(b, a [][]int) []int {
	// 将a的每一行前后相减加入前缀树
	for _, val := range a {
		Insert(splicing(val))
	}
	res := make([]int, 0, len(b))
	// 对比b数组
	for _, val := range b {
		res = append(res, countConsistent(splicing(val)))
	}
	return res
}

// 用于拼接字符串,前后相减的每个值以#分割，防止数字之间混淆
//
//	[1,2,3,4] -> "1#1#1#"
func splicing(nums []int) string {
	res := ""
	for i := 1; i < len(nums); i++ {
		num := nums[i] - nums[i-1]
		res += strconv.Itoa(num)
		res += "#"
	}
	return res
}
