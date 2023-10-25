package _go

// 静态数组做栈
const stack = [1000001]int{}

// 栈长度
var stackLen int

func dailyTemperatures(temperatures []int) []int {
	stackLen = 0
	res := make([]int, len(temperatures))
	// 遍历数组，建立小压大的单调栈
	for i := 0; i < len(temperatures); i++ {
		// 栈不为空，且栈顶元素大于当前元素
		for stackLen > 0 && temperatures[stack[stackLen-1]] > temperatures[i] {
			// 栈顶元素出栈，并记录出栈元素的下标
			top := stack[stackLen-1]
			stackLen--
			// 记录出栈元素的下标对应的数组元素
			res[top] = i - top
		}
		stack[stackLen] = i
		stackLen++
	}
	// 不用清理栈，默认为0
	return res
}
