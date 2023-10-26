package _go

// stack is a stack of integers
var stack = [151]int{}

// stackLen is the length of the stack
var stackLen int

// heights is an array of integers
var heights []int

// left, right, cur, quantity, lv, rv, le are variables used in the calculation
var left, right, cur, quantity, lv, rv, le int

// numSubmat is a function that takes in a 2D array of integers and returns an integer
func numSubmat(mat [][]int) int {
	// ans is the answer
	ans := 0
	// l is the length of the first row of the matrix
	l := len(mat[0])
	// heights is an array of integers with the length of l
	heights = make([]int, l)
	// loop through each row of the matrix
	for _, m := range mat {
		// stackLen is the length of the stack
		stackLen = 0
		// loop through each element of the row
		for j, n := range m {
			// if the element is 0, set the height to 0
			if n == 0 {
				heights[j] = 0
				// otherwise, add 1 to the height
			} else {
				heights[j] += 1
			}
			// loop through the stack until the height is less than or equal to the height of the element
			for stackLen > 0 && heights[j] <= heights[stack[stackLen-1]] {
				// cur is the index of the element in the stack
				cur = stack[stackLen-1]
				// stackLen is decremented
				stackLen--
				// left is set to -1 and lv is set to 0
				left = -1
				lv = 0
				// if stackLen is greater than 0, set left to the index of the element in the stack and lv to the height of the element in the stack
				if stackLen > 0 {
					left = stack[stackLen-1]
					lv = heights[stack[stackLen-1]]
				}
				// right is set to the index of the element
				right = j
				// rv is set to the height of the element
				rv = heights[j]
				// le is set to the difference between right and left and 1
				le = right - left - 1
				// quantity is set to the product of the height of the element and the difference between le and rv and the quantity of the rectangle
				quantity = (heights[cur] - max(lv, rv)) * le * (le + 1) / 2
				// add the quantity to the answer
				ans += quantity
			}
			// add the index of the element to the stack
			stack[stackLen] = j
			// stackLen is incremented
			stackLen++
		}
		// loop through the stack until stackLen is 0
		for stackLen > 0 {
			// cur is the index of the element in the stack
			cur = stack[stackLen-1]
			// stackLen is decremented
			stackLen--
			// left is set to -1 and lv is set to 0
			left = -1
			lv = 0
			// if stackLen is greater than 0, set left to the index of the element in the stack and lv to the height of the element in the stack
			if stackLen > 0 {
				left = stack[stackLen-1]
				lv = heights[stack[stackLen-1]]
			}
			// right is set to the length of the matrix
			right = l
			// le is set to the difference between right and left and 1
			le = right - left - 1
			// quantity is set to the product of the height of the element and the difference between le and rv and the quantity of the rectangle
			quantity = (heights[cur] - lv) * le * (le + 1) / 2
			// add the quantity to the answer
			ans += quantity
		}

	}
	// return the answer
	return ans
}

// max is a function that takes in two integers and returns the larger of the two
func max(a, b int) int {
	// if a is greater than b, return a
	if a > b {
		return a
	}
	// otherwise, return b
	return b
}
