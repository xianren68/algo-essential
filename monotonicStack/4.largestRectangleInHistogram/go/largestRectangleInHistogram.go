package _go

// 单调栈
var stack = [100001]int{}

// 栈长度
var stackLen = 0

func largestRectangleArea(heights []int) int {
	stackLen = 0
	length := len(heights)
	ans := 0
	for i := 0; i < length; i++ {
		// 找出两边的小值所在的位置得到矩形的宽，乘以当前值的高度，即为当前直方图两边扩张得到的最大矩形
		// 相同的值不用处理，因为后面相同的永远大于前面相同的值
		for stackLen > 0 || heights[i] <= heights[stack[stackLen-1]] {
			top := stack[stackLen-1]
			stackLen--
			// 左边界
			left := -1
			if stackLen > 0 {
				left = stack[stackLen-1]
			}
			// 当前面积
			area := (i - left - 1) * heights[top]
			ans = max(ans, area)

		}
		stack[stackLen] = i
		stackLen++
	}
	// 清理栈中剩余元素
	for stackLen > 0 {
		top := stack[stackLen-1]
		stackLen--
		left := -1
		if stackLen > 0 {
			left = stack[stackLen-1]
		}
		area := (length - left - 1) * heights[top]
		ans = max(ans, area)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
