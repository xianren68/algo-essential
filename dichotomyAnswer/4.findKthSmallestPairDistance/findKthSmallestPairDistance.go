package __findKthSmallestPairDistance

import "sort"

// 第k小的数，它的取值范围必然在 0~最大值-最小值中间
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	l, r := nums[0], nums[len(nums)-1]
	var mid int
	for l <= r {
		mid = l + (r-l)>>1
		if f(nums, mid) >= k {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}
func f(nums []int, m int) int {
	ans := 0
	// 滑动窗口求值
	l := 0
	r := 0
	for ; l < len(nums); l++ {
		for r+1 < len(nums) && nums[r+1]-nums[l] <= m {
			// 左指针寻找能到达的最大可以满足条件的值
			r++
		}
		ans += r - l
	}
	return ans
}
