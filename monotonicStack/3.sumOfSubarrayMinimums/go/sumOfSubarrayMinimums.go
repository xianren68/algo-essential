package _go

const MOD = 1000000007

var stack = [100000]int{}
var r int

func sumSubarrayMins(arr []int) int {
	ans := 0
	r = 0
	n := len(arr)
	for i := 0; i < n; i++ {
		for r > 0 && arr[stack[r-1]] >= arr[i] {
			top := stack[r-1]
			r--
			// 判断子数组的个数
			left := -1
			if r > 0 {
				left = stack[r-1]
			}
			ans = (ans + (i-top)*(top-left)*arr[top]) % MOD

		}
		stack[r] = i
		r++
	}
	// 栈中剩余的值
	for r > 0 {
		top := stack[r-1]
		r--
		// 判断子数组的个数
		left := -1
		if r > 0 {
			left = stack[r-1]
		}
		ans = (ans + (n-top)*(top-left)*arr[top]) % MOD
	}
	return ans
}
