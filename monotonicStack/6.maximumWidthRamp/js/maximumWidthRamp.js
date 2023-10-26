/**
 * @param {number[]} nums
 * @return {number}
 */
var maxWidthRamp = function(nums) {
    // 定义一个栈
    const stack = []
    // 获取数组的长度
    const n = nums.length
    // 定义一个变量ans，用于存储最大宽度
    let ans = 0
    // 遍历数组
    for(let i in nums){
        // 如果栈不为空，并且当前元素大于等于栈顶元素，则跳过本次循环
        if (stack.length > 0 && nums[i] >=nums[stack.at(-1)]){
            continue
        }
        // 将当前元素压入栈中
        stack.push(i)
    }
    // 遍历数组，从后往前遍历
    for(let i = n-1;i>=0 && stack.length > 0;i--){
        // 如果栈不为空，并且当前元素大于等于栈顶元素，则将栈顶元素弹出，并计算当前元素和栈顶元素之间的宽度，取最大值
        while(stack.length > 0 && nums[i]>= nums[stack.at(-1)]){
            ans = Math.max(ans,i-stack.pop())
        }
    }
    // 返回最大宽度
    return ans
};