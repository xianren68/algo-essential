/**
 * @param {character[][]} matrix
 * @return {number}
 */
var maximalRectangle = function (matrix) {
    // 定义一个栈
    const stack = []
    // 获取矩阵第一行的长度
    const n = matrix[0].length
    // 定义一个高度数组，初始值都为0
    const heights = new Array(n).fill(0)
    // 定义一个变量ans，用于存储最大矩形的面积
    let ans = 0
    // 遍历矩阵中的每一个元素
    for (let val of matrix) {
        // 遍历每一个元素中的每一个字符
        for (let i in val) {
            // 压缩数组
            // 如果当前字符为0，则高度为0，否则高度为高度数组中对应位置的高度+1
            heights[i] = val[i] == "0" ? 0 : heights[i] + 1
            // 如果栈不为空，并且栈顶元素的高度大于当前高度，则表示当前高度无法构成矩形
            while (stack.length > 0 && heights[stack.at(-1)] >= heights[i]) {
                // 将栈顶元素弹出
                let top = stack.pop()
                // 如果栈为空，则left为-1，否则left为栈顶元素
                let left = stack.length == 0 ? -1 : stack.at(-1)
                // 计算当前高度的矩形的面积，并将其与ans比较，取最大值
                let area = (i - left - 1) * heights[top]
                ans = Math.max(ans, area)
            }
            // 将当前高度压入栈中
            stack.push(i)
        }
        // 如果栈不为空，则表示栈中的元素无法构成矩形
        while (stack.length > 0) {
            // 将栈顶元素弹出
            let top = stack.pop()
            // 如果栈为空，则left为-1，否则left为栈顶元素
            let left = stack.length == 0 ? -1 : stack.at(-1)
            // 计算当前高度的矩形的面积，并将其与ans比较，取最大值
            let area = (n - left - 1) * heights[top]
            ans = Math.max(ans, area)
        }
    }
    // 返回最大矩形的面积
    return ans
};