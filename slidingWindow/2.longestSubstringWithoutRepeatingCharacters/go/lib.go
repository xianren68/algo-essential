var chars = [255]int{}

// This function returns the length of the longest substring of a given string
func lengthOfLongestSubstring(s string) int {
	// If the length of the string is 0, return 0
	if len(s) == 0 {
		return 0
	}
	// Initialize the hash table with all values set to -1
	initHash(&chars)
	// Initialize left and right pointers
	left, right := 0, 0
	// Initialize the result to 1
	res := 1
	// Iterate through the string
	for right < len(s) {
		// Set the left pointer to the maximum of the current left pointer or the index of the character in the hash table plus 1
		left = max(left, chars[s[right]]+1)
		// Set the index of the character in the hash table to the current right pointer
		chars[s[right]] = right
		// Set the result to the maximum of the current result or the difference between the right and left pointers plus 1
		res = max(res, right-left+1)
		// Increment the right pointer
		right++
	}
	// Return the result
	return res

}

// This function returns the maximum of two integers
func max(a, b int) int {
	// If a is greater than b, return a
	if a > b {
		return a
	}
	// Otherwise, return b
	return b
}

// This function initializes the hash table with all values set to -1
func initHash(chars *[255]int) {
	// Iterate through the hash table
	for i := 0; i < 255; i++ {
		// Set the value of the current index to -1
		chars[i] = -1
	}
}