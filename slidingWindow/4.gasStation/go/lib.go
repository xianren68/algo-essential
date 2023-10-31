// This function takes in two slices of integers, gas and cost, and returns an integer
func canCompleteCircuit(gas []int, cost []int) int {
	// Get the length of the cost slice
	n := len(cost)
	// Initialize left and right indices
	for l, r := 0, 0; l < n; {
		// Initialize length and sum
		leng, sum := 0, 0
		// Iterate until the sum is negative
		for sum >= 0 {
			// If the length is equal to the length of the cost slice, return the left index
			if leng == n {
				return l
			}
			// Calculate the right index
			r = (l + leng) % n
			// Increment the length
			leng++
			// Add the difference between the gas and cost slices at the left and right indices
			sum += gas[r] - cost[r]
		}
		// Increment the left index
		l += leng
	}
	// If the loop is finished without returning, return -1
	return -1
}