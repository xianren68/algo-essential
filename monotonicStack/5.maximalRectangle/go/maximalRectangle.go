package _go

// 本题与上一题差不多，我们只需要以每一行为底，计算最大矩形的面积即可
// 每次遍历完一行，用下一行累加当前行的值，压缩数组（当前行值为0,则不压缩）
// 栈
var stack = [201]int{}

// 栈长度
var stackLen int

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix[0])
	// 记录每一行的高度
	heights := make([]int, n)
	ans := 0
	for _, mat := range matrix {
		stackLen = 0
		for i := range mat {
			// 压缩数组
			if mat[i] == '0' {
				heights[i] = 0
			} else {
				heights[i] = heights[i] + 1
			}
			// 栈不为空，且当前高度大于栈顶高度时，弹出栈顶元素，计算面积，并继续弹出栈顶元素
			for stackLen > 0 && heights[stack[stackLen-1]] >= heights[i] {
				top := stack[stackLen-1]
				stackLen--
				left := -1
				if stackLen > 0 {
					left = stack[stackLen-1]
				}
				area := (i - left - 1) * heights[top]
				ans = max(ans, area)
			}
			stack[stackLen] = i
			stackLen++
		}
		// 清理栈
		for stackLen > 0 {
			top := stack[stackLen-1]
			stackLen--
			left := -1
			if stackLen > 0 {
				left = stack[stackLen-1]
			}
			area := (n - left - 1) * heights[top]
			ans = max(ans, area)
		}

	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
