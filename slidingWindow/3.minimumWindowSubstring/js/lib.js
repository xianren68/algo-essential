/**
 * @param {string} s
 * @param {string} t
 * @return {string}
 */
const cnts = new Array(256)
var minWindow = function(s, t) {
    // 如果s的长度小于t的长度，直接返回空字符串
    if(s.length < t.length) return ''
    // 将cnts数组填充为0
    cnts.fill(0)
    // 遍历t，将cnts数组中对应的字符的值减1
    for(let i of t){
        cnts[i.charCodeAt(0)]--
    }
    // 定义左右指针，minSize用来记录最小窗口大小，minStart用来记录最小窗口的起始位置
    let left = 0, right = 0
    let minSize = Number.MAX_VALUE
    let minStart = 0
    let all = t.length
    // 遍历s，如果cnts数组中对应的字符的值大于0，则将all减1
    for(;right < s.length; right++){
        if (cnts[s[right].charCodeAt(0)]++ < 0){
            all --
        }
        // 如果all等于0，则说明窗口中包含t中的所有字符
        if(all === 0){
            // 将cnts数组中对应的字符的值加1，并将left指针右移，直到cnts数组中对应的字符的值大于0
            while(cnts[s[left].charCodeAt(0)] > 0){
                cnts[s[left++].charCodeAt(0)]--
            }
            // 如果right-left+1小于minSize，则更新minSize和minStart的值
            if (right - left + 1 < minSize){
                minSize = right - left + 1
                minStart = left
            }
        }
    }
    // 如果minSize的值没有更新，则说明不存在符合条件的窗口，返回空字符串，否则返回s中从minStart到minStart+minSize的字符串
    return minSize === Number.MAX_VALUE ? '' : s.slice(minStart, minStart + minSize)
};