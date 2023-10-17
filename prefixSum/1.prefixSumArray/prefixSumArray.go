package __prefixSumArray

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	// 创建一个前缀和数组
	// 记录第零位前缀和为0,避免判断边界条件
	prefixSum := make([]int, len(nums)+1)
	sum := 0
	for i, val := range nums {
		prefixSum[i] = sum
		sum += val
	}
	prefixSum[len(nums)] = sum
	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}
