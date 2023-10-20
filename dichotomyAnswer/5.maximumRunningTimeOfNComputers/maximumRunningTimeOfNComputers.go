package __maximumRunningTimeOfNComputers

// 通过分析我们可以知道两个结论
// 1. 如果有电池电量超过需要的时间，则此电池负责一台电脑即可
// 2. 如果全是碎片电池，没有超过需要时间的电池，则 （碎片电池总电量 >= 电脑数量*时间 ）时所有电脑都可以开启此时间
func maxRunTime(n int, batteries []int) int64 {
	// 所有电池电量的总和
	sum := 0
	max := batteries[0]
	for _, battery := range batteries {
		if max < battery {
			max = battery
		}
		sum += battery
	}
	for sum > max*n {
		// 说明 max < 需要的时间，即集合中全是碎片电池
		return int64(sum / n)
	}
	l, r := 0, max
	res := 0
	var mid int
	for l <= r {
		mid = l + (r-l)>>1
		if f(batteries, mid, sum, n) {
			// 满足条件，看有无更大的
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return int64(res)

}

// 判断给定time时间，所有电脑能否同时运行
func f(nums []int, time, sum, n int) bool {
	for _, num := range nums {
		if num > time {
			// 负责一块电池
			n--
			sum -= num
		}
	}
	return sum >= n*time

}
