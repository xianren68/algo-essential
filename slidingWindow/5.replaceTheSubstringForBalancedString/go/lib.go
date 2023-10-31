// This function takes a string as an argument and returns an integer
func balancedString(s string) int {
	// Initialize variables to store the length of the string and the counts of each character
	n := len(s)
	cnts := [4]int{}
	// Create an array to store the index of each character
	arr := make([]int, n)
	// Iterate through the string and store the index of each character in the array
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'Q':
			arr[i] = 0
			cnts[0]++
		case 'W':
			arr[i] = 1
			cnts[1]++
		case 'E':
			arr[i] = 2
			cnts[2]++
		case 'R':
			arr[i] = 3
			cnts[3]++
		}
	}
	// Calculate the target number of each character
	target := n / 4
	// Initialize the answer to the length of the string
	ans := n
	// Check if the counts of each character are equal to the target number
	if ok(cnts[:], 0, target) {
		return 0
	}
	// Iterate through the string and check if the counts of each character are equal to the target number
	for r, l := 0, 0; r < n; r++ {
		cnts[arr[r]]--
		// If the counts of each character are equal to the target number, store the length of the substring
		for ok(cnts[:], r-l+1, target) {
			ans = min(ans, r-l+1)
			// Restore the counts of each character
			cnts[arr[l]]++
			l++
		}
	}
	// Return the length of the substring
	return ans
}

// This function checks if the counts of each character are equal to the target number
func ok(cnts []int, leng, target int) bool {
	// Iterate through the counts of each character
	for i := 0; i < 4; i++ {
		// If the counts of each character are greater than the target number, return false
		if cnts[i] > target {
			return false
		}
		// Subtract the target number from the length of the substring
		leng -= target - cnts[i]
	}
	// If the length of the substring is 0, return true
	return leng == 0
}

// This function returns the minimum of two integers
func min(a, b int) int {
	// If a is greater than b, return b
	if a > b {
		return b
	}
	// Otherwise, return a
	return a
}