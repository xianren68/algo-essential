package _go

// This function removes duplicate letters from a given string
func removeDuplicateLetters(s string) string {
	// Create an array of integers to store the quantity of each letter
	quantity := make([]int, 26)
	// Create a stack to store the letters
	stack := [26]byte{}
	// Initialize the stack length to 0
	stackLen := 0
	// Create an array of booleans to store whether a letter has been used
	used := make([]bool, 26)
	// Iterate through the given string
	for i := 0; i < len(s); i++ {
		// Increment the quantity of each letter
		quantity[s[i]-'a']++
	}
	// Iterate through the given string
	for i := 0; i < len(s); i++ {
		// If the letter has already been used, increment the quantity
		if used[s[i]-'a'] {
			quantity[s[i]-'a']--
			continue
		}
		// If the letter is not the last letter in the stack, and the letter is less than the last letter in the stack, and the quantity of the letter is greater than 0
		for stackLen > 0 && s[i] < stack[stackLen-1] && quantity[stack[stackLen-1]-'a'] > 0 {
			// Set the used boolean to false
			used[stack[stackLen-1]-'a'] = false
			// Decrement the stack length
			stackLen--
		}
		// Push the letter onto the stack
		stack[stackLen] = s[i]
		// Increment the stack length
		stackLen++
		// Set the used boolean to true
		used[s[i]-'a'] = true
		// Decrement the quantity
		quantity[s[i]-'a']--
	}
	// Return the stack as a string
	return string(stack[:stackLen])
}
