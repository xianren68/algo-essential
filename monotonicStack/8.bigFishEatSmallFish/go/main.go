package main

import (
	"bufio"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func read() int {
	flag := false
	ret := 0
	for c, _ := in.ReadByte(); (c >= '0' && c <= '9') || c == '-'; c, _ = in.ReadByte() {
		if c == '-' {
			flag = true
			continue
		}
		ret = (ret * 10) + int(c-'0')
	}
	if flag {
		ret = -ret
	}
	return ret
}
func main() {
	defer out.Flush()
	n := read()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = read()
	}
	ans := compute(nums)
	out.WriteString(strconv.Itoa(ans))
	return
}
func compute(nums []int) int {
	n := len(nums)
	stack := make([][2]int, n)
	stackLen := 0
	ans := 0
	for i := n - 1; i >= 0; i-- {
		cur := 0
		// 将前面的小鱼全部吃掉需要多少轮
		for stackLen > 0 && nums[i] > stack[stackLen-1][0] {
			// 吃掉小鱼
			cur = max(cur+1, stack[stackLen-1][1])
			stackLen--
		}
		stack[stackLen] = [2]int{nums[i], cur}
		stackLen++
		ans = max(ans, cur)
	}
	return ans

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
