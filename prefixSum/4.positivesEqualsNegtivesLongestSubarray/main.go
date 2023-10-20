package main

import (
	"bufio"
	"os"
	"strconv"
)

// 本题与计算累加和为指定数据的子数组最长长度类似

// 将正 -> 1,负 -> -1
// 这样正负数相等时累加和即为0，我们将目标值设为0即可

const MAX = 100001

// 数字数组
var arr = [MAX]int{}
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

var m int

// 不用清理静态数组，因为我们每次使用的地方会覆盖
func main() {
	m = read()
	for i := 0; i < m; i++ {
		arr[i] = read()
	}
	res := handler(arr[:m])
	out.WriteString(strconv.Itoa(res))
	out.Flush()
}
func handler(nums []int) int {
	// 创建前缀hash表，表中只记录每个前缀和出现的最早位置
	// 前缀和为0的时候，在-1位置
	prefixSum := make(map[int]int, 20)
	prefixSum[0] = -1
	sum := 0
	max := 0
	// 计算每个位置和目标的差值
	for i, val := range nums {
		// 不用建一个新数组来改变值，在累加时改变
		if val > 0 {
			sum += 1
		} else if val < 0 {
			sum -= 1
		} else {
			sum += val
		}
		want := sum - 0
		// 判断hash表中是否存在
		if index, ok := prefixSum[want]; ok {
			// 已存在，记录以当前位置为边界的子数组的最长长度
			l := i - index
			if l > max {
				max = l
			}
		}
		// 判断当前前缀和是否已经存在于hash表中
		if _, ok := prefixSum[sum]; !ok {
			// 存进去
			prefixSum[sum] = i
		}
	}
	return max
}
