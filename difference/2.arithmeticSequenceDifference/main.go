package main

import (
	"bufio"
	"os"
	"strconv"
)

var nums = [10000010]int{}
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

// 在一个范围内累加一个等差数列
// 以2 - 10 为例，如果每一项都为等差数，则求前缀和可得结果[2,2,2,2,2] -> [2,4,6,8,10]
// 而要使它们都为同一个值，则与上题的步骤一致
// 所以经历两次累加即可
func set(l, r, s, e int) {
	// 等差数列每一项的差
	item := (e - s) / (r - l)
	nums[l] += s
	nums[l+1] += item - s
	nums[r+1] -= item + e
	nums[r+2] += e
}

// 两次前缀和
func build(n int) {
	for i := 1; i <= n; i++ {
		nums[i] += nums[i-1]
	}
	for i := 1; i <= n; i++ {
		nums[i] += nums[i-1]
	}

}

func main() {
	n, m := read(), read()
	for i := 0; i < m; i++ {
		l, r, s, e := read(), read(), read(), read()
		set(l, r, s, e)
	}
	build(n)
	max, xor := 0, 0
	for i := 1; i <= n; i++ {
		if max < nums[i] {
			max = nums[i]
		}
		xor ^= nums[i]
	}
	out.WriteString(strconv.Itoa(xor) + " " + strconv.Itoa(max))
	out.Flush()
}
