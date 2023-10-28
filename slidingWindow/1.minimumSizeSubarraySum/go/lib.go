// Function to return the minimum length subarray with sum greater than or equal to the target
func minSubArrayLen(target int, nums []int) int {
	// Initialize left and right pointers
	left, right := 0, 0
	// Initialize sum to 0
	sum := 0
	// Get the length of the array
	l := len(nums)
	// Initialize result to the length of the array + 1
	// 设置为比数组长度+1，便于判断整个数组是否有 > target的子数组
	res := l + 1
	// Iterate until the right pointer is less than the length of the array
	for ; right < l; right++ {
		// Add the current element to the sum
		sum += nums[right]
		// Iterate until the sum is greater than or equal to the target
		for sum >= target {
			// Update the result if the current subarray length is less than the current result
			res = min(res, right-left+1)
			// Subtract the left element from the sum
			sum -= nums[left]
			// Increment the left pointer
			left++
		}
	}
	// If the result is still the length of the array + 1, return 0
	if res == l+1 {
		return 0
	}
	// Return the result
	return res
}

// Function to return the minimum of two integers
func min(a, b int) int {
	// Return the smaller of the two integers
	if a > b {
		return b
	}
	return a
}