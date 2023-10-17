package main

import "fmt"

var trie = [3000001][2]int{}
var cnt = 1

// HighNum 找到最大值，找到它的最高位，算出它前缀0的个数，我们从不是零的位开始建立前缀树
func HighNum(nums []int) int {
	max := nums[0]
	for _, val := range nums {
		if max < val {
			max = val
		}
	}
	if max == 0 {
		return -1
	}
	res := 0
	for i := 0; i < 32; i++ {
		if max>>i == 1 {
			res = i
		}
	}
	return res
}

func Insert(high int, num int) {
	cur := 1
	for i := high; i >= 0; i-- {
		path := (num >> i) & 1
		if trie[cur][path] == 0 {
			trie[cur][path] = cnt + 1
			cnt++
		}

		cur = trie[cur][path]
	}
}

// SearchMax 在每一位上寻找能让异或值更大的数
func SearchMax(high int, num int) int {
	cur := 1
	res := 0
	for i := high; i >= 0; i-- {

		// 当前位的值
		status := (num >> i) & 1

		// 想要的值
		want := status ^ 1
		// 查看是否存在想要的路径
		if trie[cur][want] == 0 {
			// 不存在，还是原值
			want ^= 1
		}
		// 将这一位最终的值加入
		res |= (status ^ want) << i
		cur = trie[cur][want]
	}
	return res
}

func findMaximumXOR(nums []int) int {
	cnt = 1
	max := 0
	high := HighNum(nums)
	// 建立树
	for _, val := range nums {
		Insert(high, val)
	}
	// 查找最大值
	for _, val := range nums {
		res := SearchMax(high, val)
		if res > max {
			max = res
		}

	}
	clear()
	return max

}
func clear() {
	for i := 1; i <= cnt; i++ {
		trie[i] = [2]int{}
	}
}
func main() {
	fmt.Println(findMaximumXOR([]int{2, 4}))
}
