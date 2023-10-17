package longestWellPerformingInterval

func longestWPI(hours []int) int {
	// 前缀和
	sum := 0
	// 记录某个前缀最早出现的次数
	hash := make(map[int]int)
	hash[0] = -1
	res := 0
	// 判断以每个日期为边界的最长良好工作时间
	for i, hour := range hours {
		// 8小时以上为1，以下为-1
		if hour > 8 {
			sum += 1
		} else {
			sum -= 1
		}
		if sum > 0 {
			if i+1 > res {
				res = i + 1
			}
		} else if index, ok := hash[sum-1]; ok { // 前缀和为1，即表示劳累的天数多
			// 存在，记录当前位置为底的最长长度
			if i-index > res {
				res = i - index
			}
		}
		// 判断当前前缀和在hash表中是否存在
		if _, ok := hash[sum]; !ok {
			// 不存在，添加
			hash[sum] = i
		}
	}
	return res
}
