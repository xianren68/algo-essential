package main

import (
	"bufio"
	"os"
	"strconv"
)

// 初始能量最大与最小区间  所有建筑最大值-0
// 单调性 初始能量越多，越容易过关

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func read() int {
	// 指示是否有-号
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
	n := read()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = read()
	}
	_, _ = out.WriteString(strconv.Itoa(process(nums)))
	_ = out.Flush()
}
func process(nums []int) int {
	ans := 0
	// 找出最大值
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	// 0 - max不断二分
	l, r := 0, max
	for l <= r {
		mid := l + (r-l)>>1
		if f(mid, max, nums) {
			// 记录答案
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}

// 判断机器人在num能量下能否过关
func f(num, max int, nums []int) bool {
	for _, val := range nums {
		if num >= val {
			num += num - val
		} else {
			num -= val - num
		}
		// 能量已经达到所有建筑的最高值了，肯定能过关
		if num >= max {
			return true
		} else if num < 0 {
			// 能量为负
			return false
		}
	}
	return true
}
