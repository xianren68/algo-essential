/**
 * @param {number} target
 * @param {number[]} nums
 * @return {number}
 */
var minSubArrayLen = function(target, nums) {
    // 定义左右指针
    let left = 0
    let right = 0
    // 定义累加和
    let sum = 0
    // 获取数组长度
    let l = nums.length
    // 初始化结果
    let res = l + 1
    // 当右指针小于数组长度时，循环
    while (right < l) {
        // 累加
        sum += nums[right]
        // 当累加和大于等于目标值时，更新结果，并且将累加和减去左指针指向的值
        while (sum >= target) {
            res = Math.min(res, right - left + 1)
            sum -= nums[left]
            left++
        }
        // 右指针右移
        right++
    }
    // 如果结果没有变化，则返回0，否则返回结果
    return res === l + 1 ? 0 : res
    
};