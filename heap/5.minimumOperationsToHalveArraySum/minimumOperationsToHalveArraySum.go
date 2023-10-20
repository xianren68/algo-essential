package __minimumOperationsToHalveArraySum

var heap []float64
var size int

// 初始化堆
func initHeap(n int) {
	size = 0
	heap = make([]float64, n)
}

// 添加
func push(n float64) {
	heap[size] = n
	i := size
	size++
	// 向上替换
	for i > 0 && heap[(i-1)/2] < heap[i] {
		heap[(i-1)/2], heap[i] = heap[i], heap[(i-1)/2]
		i = (i - 1) / 2
	}
}

// 弹出
func pop() float64 {
	ret := heap[0]
	heap[0] = heap[size-1]
	size--
	i := 0
	l := i*2 + 1
	var best int
	for l < size {
		if l+1 < size && heap[l] < heap[l+1] {
			best = l + 1
		} else {
			best = l
		}
		if heap[i] < heap[best] {
			heap[i], heap[best] = heap[best], heap[i]
			i = best
			l = i*2 + 1
		} else {
			break
		}
	}
	return ret
}
func halveArray(nums []int) int {
	// 构建大根堆
	initHeap(len(nums))
	var sum float64
	for _, num := range nums {
		sum += float64(num)
		push(float64(num))
	}
	ans := 0
	curSum := sum
	for curSum > sum/2 {
		// 弹出最大值，除二
		num := pop()
		num = num / 2
		curSum -= num
		push(num)
		ans++
	}
	return ans
}
