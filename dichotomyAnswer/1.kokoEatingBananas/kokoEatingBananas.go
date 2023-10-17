package __kokoEatingBananas

func minEatingSpeed(piles []int, h int) int {
	// 最小速度
	l := 1
	// 最大的达标速度
	r := 0
	res := 0
	// 找到所有需要项中所需时间最多的，作为达标的最大速度
	for _, val := range piles {
		if val > r {
			r = val
		}
	}
	// 在l-r不断二分，找到最小且达标的速度
	for l <= r {
		mid := l + (r-l)>>1
		if f(piles, mid, h) {
			// 达标，记录答案，查看是否有无更小的
			res = mid
			r = mid - 1
		} else {
			// 不达标
			l = mid + 1
		}
	}
	return res
}

// 判断当前速度是否达标
func f(piles []int, speed int, h int) bool {
	// 按照当前速度，吃完香蕉的时间
	sumH := 0
	for _, pile := range piles {
		// 向上取整
		sumH += (pile + speed - 1) / speed
	}
	return sumH <= h
}
