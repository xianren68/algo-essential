/**
 * @param {string} s
 * @return {string}
 */
var removeDuplicateLetters = function (s) {
    // 创建一个栈
    const stack = []
    // 创建一个数组，用来记录每个字母出现的次数
    const quantity = new Array(26).fill(0)
    // 创建一个数组，用来记录每个字母是否被使用
    const used = new Array(26).fill(false)
    // 遍历字符串，记录每个字母出现的次数
    for (let i of s) {
        quantity[i.charCodeAt(0) - 97]++
    }
    // 遍历字符串
    for (let i of s) {
        // 如果该字母已被使用，则跳过
        if (used[i.charCodeAt(0) - 97]) {
            quantity[i.charCodeAt(0) - 97]--
            continue
        }
        // 如果栈不为空，且栈顶元素大于当前元素，且栈顶元素在字符串中仍存在，则弹出栈顶元素
        while (stack.length > 0 && stack.at(-1) > i && quantity[stack.at(-1).charCodeAt(0) - 97] > 0) {
            used[stack.at(-1).charCodeAt(0) - 97] = false
            stack.pop()
        }
        // 将当前元素压入栈中
        stack.push(i)
        // 标记当前元素已被使用
        used[i.charCodeAt(0) - 97] = true
        // 将当前元素出现的次数减一
        quantity[i.charCodeAt(0) - 97]--
    }
    // 返回栈中的元素组成的字符串
    return stack.join("")
};