package __evenCountsLongestSubarray

// 记录每个位置每个元音出现的次数的奇偶性
// 然后以每个位置为底边寻找最长子串
// 与当前位置元音奇偶性相同的位置即表示从那个位置到当前位置为偶数个，取最早出现的，即为最长的
func findTheLongestSubstring(s string) int {
	var res = 0
	// 只需要五位来表示每个元音字母的奇偶性
	var evenMap = map[byte]int{0: -1}
	// 0..i的奇偶性
	var evenPre byte = 0
	for i := 0; i < len(s); i++ {
		b := mapEven(s[i])
		if b >= 0 {
			// 修改奇偶性
			evenPre ^= 1 << b
		}
		// 寻找相同的
		if index, ok := evenMap[evenPre]; ok {
			if res < i-index {
				res = i - index
			}
			continue
		}
		evenMap[evenPre] = i
	}
	return res
}

// 映射元音位
func mapEven(val byte) int {
	switch val {
	case 'a':
		return 0
	case 'e':
		return 1
	case 'i':
		return 2
	case 'o':
		return 3
	case 'u':
		return 4
	default:
		return -1
	}
}
