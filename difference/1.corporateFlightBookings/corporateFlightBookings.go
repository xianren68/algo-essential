package __corporateFlightBookings

func corpFlightBookings(bookings [][]int, n int) []int {
	// 多一个长度，避免边界判断
	res := make([]int, n+2)
	for _, book := range bookings {
		res[book[0]] += book[2]   // 起始位置加订阅数量
		res[book[1]+1] -= book[2] // 结束位置减去订阅数量
	}
	for i := 1; i <= n; i++ {
		res[i] += res[i-1]
	}
	return res[1 : n+1]

}
