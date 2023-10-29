// This function returns the minimum window of a string s and a string t.
func minWindow(s string, t string) string {
	// If the length of s is less than the length of t, return an empty string.
	if len(s) < len(t) {
		return ""
	}
	// Create an array of integers to store the number of times each character appears in t.
	cnts := [256]int{}
	for i := 0; i < len(t); i++ {
		cnts[t[i]]--
	}
	// Initialize the minimum size and start index of the window.
	minSize := math.MaxInt
	start := 0
	all := len(t)
	// Iterate through the characters of s.
	for right, left := 0, 0; right < len(t); right++ {
		// If the number of times the character appears in t is less than 0, increment all.
		if cnts[s[right]] < 0 {
			all--
		}
		cnts[s[right]]++
		// If all is 0, the window is valid.
		if all == 0 {
			// Iterate through the characters of s from left to right until the number of times the character appears in t is greater than 0.
			for cnts[s[left]] > 0 {
				cnts[s[left]]--
				left++
			}
			// If the size of the window is less than the current minimum size, update the minimum size and start index.
			if minSize > right-left+1 {
				minSize = right - left + 1
				start = left
			}
		}
	}
	// If the minimum size is still math.MaxInt, return an empty string.
	if minSize == math.MaxInt {
		return ""
	}
	// Return the substring of s from start to start+minSize.
	return s[start : start+minSize]
}