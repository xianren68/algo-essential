/**
 * @param {number[]} gas
 * @param {number[]} cost
 * @return {number}
 */
var canCompleteCircuit = function(gas, cost) {
    const n = gas.length;
    // 从0开始，尝试每一个位置作为起点
    for (let l = 0, r = 0; l < n; ) {
        // 记录当前路径的长度
        let leng = 0
        // 记录当前路径的总和
        let sum = 0
        // 当总和大于等于0时，说明当前路径是可行的
        while(sum >= 0){
            // 如果当前路径的长度等于数组的长度，说明找到了可行的路径
            if(leng == n){
                return l
            }
            // 尝试下一个位置作为当前路径的起点
            r = (l + leng)%n
            leng++
            // 更新当前路径的总和
            sum += gas[r] - cost[r]
        }
        // 尝试下一个位置作为起点
        l += leng
    }
    // 如果没有找到可行的路径，返回-1
    return -1
};