package _go

// stack is a stack of ints with a length of 100000
var stack = [100000]int{}

// stackLen is the length of the stack
var stackLen = 0

// maxWidthRamp is a function that takes in an array of ints and returns an int
func maxWidthRamp(nums []int) int {
	// set the length of the stack to 0
	stackLen = 0
	// set the answer to 0
	ans := 0
	// get the length of the array
	n := len(nums)
	// loop through the array
	for i := 0; i < n; i++ {
		// if the stack length is greater than 0 and the current number is greater than or equal to the number in the stack, continue
		if stackLen > 0 && nums[i] >= nums[stack[stackLen-1]] {
			continue
		}
		// set the current number to the current index of the stack
		stack[stackLen] = i
		// increment the stack length
		stackLen++
	}
	// loop through the array in reverse order
	for i := n - 1; i >= 0 && stackLen > 0; i-- {
		// if the stack length is greater than 0 and the current number is greater than or equal to the number in the stack, set the answer to the difference between the current number and the number in the stack, and decrement the stack length
		for stackLen > 0 && nums[i] >= nums[stack[stackLen-1]] {
			ans = max(ans, i-stack[stackLen-1])
			stackLen--
		}
	}
	// return the answer
	return ans

}

// max is a function that takes in two ints and returns the larger of the two
func max(a, b int) int {
	// if a is greater than b, return a
	if a > b {
		return a
	}
	// otherwise, return b
	return b
}
