package main

import (
	"bufio"
	"os"
	"strconv"
)

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
func compute(nums []int) [][2]int {
	// 定义一个单调栈
	stack := make([]int, 0, len(nums))
	// 一个二维数组用来存储对应值
	res := make([][2]int, len(nums))
	// 下标入栈，如果碰到大的就让它们出栈
	for i, num := range nums {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		// 将大于等于当前值的值出栈
		for len(stack) > 0 && num <= nums[stack[len(stack)-1]] {
			// 出栈
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 前面第一个小值为栈中第一个位置
			// 后面第一小值为让自己出栈的值
			if len(stack) < 1 {
				res[n] = [2]int{-1, i}
			} else {
				res[n] = [2]int{stack[len(stack)-1], i}
			}
		}
		stack = append(stack, i)
	}
	// 清理栈
	for len(stack) > 1 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res[n] = [2]int{stack[len(stack)-1], -1}
	}
	// 最后一个值
	res[stack[0]] = [2]int{-1, -1}
	// 整理，将重复值的小值归位
	for i := len(res) - 2; i >= 0; i-- {
		// 有相同值的影响
		if res[i][1] != -1 && nums[res[i][1]] == nums[i] {
			res[i][1] = res[res[i][1]][1]
		}
	}
	return res

}
func main() {
	n := read()
	nums := make([]int, n)
	var m int
	for i := 0; i < n; i++ {
		m = read()
		nums[i] = m
	}
	res := compute(nums)
	for _, re := range res {
		out.WriteString(strconv.Itoa(re[0]))
		out.WriteString(" ")
		out.WriteString(strconv.Itoa(re[1]))
		out.WriteString("\n")
	}
	out.Flush()
}
