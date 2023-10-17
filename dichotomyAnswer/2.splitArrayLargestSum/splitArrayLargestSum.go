package __splitArrayLargestSum

// 分割后最大可以满足题意的值为nums中的所有值的和，符合题意的值在0-sum之间
// 分割的值越小，则需要的k越大，我们找到分割到正好可以满足题意的值即为所需的答案
func splitArray(nums []int, k int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 在0-max之间不断二分寻找最适合的答案
	l, r := 0, sum
	res := 0
	for l <= r {
		mid := l + (r-l)>>1
		v := f(nums, mid)
		if v == -1 || v > k {
			// 划分的块太小
			l = mid + 1
		} else {
			// 合适的答案，看看有没有划分更小的答案
			res = mid
			r = mid - 1
		}
	}
	return res
}

// 返回以当前值分割所需要分割成的区域
func f(nums []int, num int) (res int) {
	res = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > num {
			// 分割的太小了
			return -1
		}
		if sum+nums[i] > num {
			sum = nums[i]
			res++
		} else {
			sum += nums[i]
		}
	}
	return
}
