
/**
 * @param {number[]} arr
 * @return {number}
 */
const MOD = 1000000007
var sumSubarrayMins = function(arr) {
    const stack = []
    const n = arr.length
    let ans = 0
    for(let i in arr){
        while(stack.length > 0 && arr[stack.at(-1)] >= arr[i]){
            const top = stack.pop()
            // 以当前值为最小值的子数组有多少个
            // 左边界
            const left = stack.length == 0 ?-1:stack.at(-1)
            ans = (ans + (top-left)*(i-top)*arr[top]) % MOD
        }
				stack.push(i)
    }
    // 清理栈中剩余的值
    while(stack.length > 0){
        const top = stack.pop()
        const left = stack.length == 0 ?-1:stack.at(-1)
        ans = (ans + (top-left)*(n-top)*arr[top]) % MOD
    }
    return ans
};