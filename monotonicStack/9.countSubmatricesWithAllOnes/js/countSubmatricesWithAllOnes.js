/**
 * @param {number[][]} mat
 * @return {number}
 */
var numSubmat = function(mat) {
    // 获取矩阵第一列的长度
    const n = mat[0].length
    // 初始化高度数组
    const heights = new Array(n).fill(0)
    // 初始化栈
    const stack = []
		let ans = 0
    // 遍历矩阵
    for(let m of mat){
        // 遍历每一行
        for(let i in m){
            // 更新高度数组
            heights[i] = m[i] == 0 ? 0:heights[i]+1
            // 当栈不为空，且当前高度小于栈顶高度时，出栈，计算子矩阵面积
            while (stack.length > 0 && heights[i] <= heights[stack.at(-1)]){
                cur = stack.pop()
                let left = stack.length == 0?-1:stack.at(-1)
                let right = i
                len = right - left -1
                ans += (heights[cur]-Math.max(left == -1?0:heights[left],heights[right]))*len*(len+1)/2
            }
            // 将当前高度入栈
            stack.push(i)
        }
        // 当栈不为空时，出栈，计算子矩阵面积
        while(stack.length>0){
            cur = stack.pop()
                let left = stack.length == 0?-1:stack.at(-1)
                let right = n
                len = right - left -1
                ans += (heights[cur]-(left == -1?0:heights[left]))*len*(len+1)/2
        }
    }
    // 返回子矩阵面积
    return ans
};