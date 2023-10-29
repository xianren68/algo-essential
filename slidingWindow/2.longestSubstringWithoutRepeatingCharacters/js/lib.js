/**
 * @param {string} s
 * @return {number}
 */
let record = new Array(255)

function lengthOfLongestSubstring(s) {
 // 如果字符串长度为0，则返回0
 if (s.length == 0) {
        return 0
    }
    // 定义左右指针
    let left = 0
    let right = 0
    // 记录最大子串长度
    let res = 1
    // 初始化记录数组，记录每个字符的索引
    record.fill(-1)
 // 当右指针小于字符串长度时，循环
 while(right < s.length){
        // 左指针大于记录索引时，更新左指针，否则不变
        left = Math.max(left, record[s[right].charCodeAt(0)] + 1)
        // 记录最大子串长度
        res = Math.max(res, right - left + 1)
        // 更新记录数组，记录每个字符的索引
        record[s[right].charCodeAt(0)] = right
        // 右指针右移
        right++
    }
    // 返回最大子串长度
    return res

}