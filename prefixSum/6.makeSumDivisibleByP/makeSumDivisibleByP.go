package makeSumDivisibleByP

func minSubarray(nums []int, p int) int {
	// 求出总和
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 求出总和余数
	mod := sum % p
	if mod == 0 {
		return 0
	}
	// 记录某一个余数出现的最晚位置
	hash := map[int]int{0: -1}
	// 前缀和
	prefix := 0
	res := len(nums)
	// 遍历数组，记录每一位的前缀和产生的余数和总余数的差距
	for i, num := range nums {
		prefix += num
		// 当前余数
		curmod := prefix % p
		// 还需要的余数
		wantmod := (curmod + p - mod) % p
		if index, ok := hash[wantmod]; ok {
			if res > i-index {
				res = i - index
			}
		}
		// 存储当前余数
		hash[curmod] = i
	}
	if res == len(nums) {
		return -1
	}
	return res
}
