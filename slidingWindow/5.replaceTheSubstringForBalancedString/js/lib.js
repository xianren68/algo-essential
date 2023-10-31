/**
 * @param {string} s
 * @return {number}
 */
var balancedString = function(s) {
    const n = s.length
    const cnts = new Array(4).fill(0)
    const arr = new Array(n).fill(0)
    for (let i = 0; i < n; i++) {
        switch (s[i]) {
            case 'Q':
                cnts[0]++
                arr[i] = 0
                break
            case 'W':
                cnts[1]++
                arr[i] = 1
                break
            case 'E':
                cnts[2]++
                arr[i] = 2
                break
            case 'R':
                cnts[3]++
                arr[i] = 3
                break
        }
    }
    let ans = n
    const target = n/4
    // 如果cnts数组中每个元素都小于target，则返回0
    if(ok(cnts, 0, target)){
        return 0
    }
    // 从左到右遍历字符串，如果cnts数组中每个元素都大于target，则更新ans，并使cnts数组中每个元素减1，leng加1
    for(let l = 0,r = 0;r < n; r++){
        cnts[arr[r]]--
        while(ok(cnts, r - l + 1, target)){
            ans = Math.min(ans, r - l + 1)
            cnts[arr[l]]++
            l++
        }
    }
    return ans
};
// 判断cnts数组中每个元素是否都大于target
function ok(cnts, leng,target){
    for(let i = 0; i < 4; i++){
        if(cnts[i] > target){
            return false
        }
        leng -= target - cnts[i]

    }
    return leng == 0
}