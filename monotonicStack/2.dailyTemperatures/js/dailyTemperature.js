/**
 * @param {number[]} temperatures
 * @return {number[]}
 */
var dailyTemperatures = function(temperatures) {
    const stack = []
    // 建立一个长度为传入数组长度的数组，全部填充0
    const res = new Array(temperatures.length).fill(0)
    for (let i in temperatures){
        while(stack.length > 0 && temperatures[stack.at(-1)] < temperatures[i]){
            let top = stack.pop()
            res[top] = i - top
        }
        stack.push(i)
    }
    return res

};