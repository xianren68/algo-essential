package numberOfSubarraySumEqualsAim

// 判断以每个数为边界时，可以达成目标的子数组有多少个
func subarraySum(nums []int, k int) int {
	// hash 表存储每个累加和出现的次数
	numberSum := make(map[int]int, 20)
	// 累加和为0时在0位置前有一次
	numberSum[0] = 1
	// 累加和
	sum := 0
	res := 0
	for _, val := range nums {
		sum += val
		want := sum - k
		res += numberSum[want]
		numberSum[sum]++
	}
	return res
}
