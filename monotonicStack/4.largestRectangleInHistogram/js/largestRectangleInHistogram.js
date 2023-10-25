/**
 * @param {number[]} heights
 * @return {number}
 */
const stack = new Array(1000001).fill(0)
let stackLen
var largestRectangleArea = function(heights) {
    let ans = 0
    stackLen = 0
    for(let i in heights){
        while (stackLen > 0 && heights[i] <= heights[stack[stackLen-1]]){
            let top = stack[--stackLen]
            let left = stackLen > 0 ? stack[stackLen-1]:-1
            let area = (i-left-1)*heights[top]
            ans = max(ans,area)
        }
        stack[stackLen++] = i
    }
    while (stackLen > 0){
        let top = stack[--stackLen]
            let left = stackLen > 0 ? stack[stackLen-1]:-1
            let area = (heights.length-left-1)*heights[top]
            ans = max(ans,area)
    }
    return ans
}
function max(a,b){
    if(a > b){
        return a
    }
    return b
}