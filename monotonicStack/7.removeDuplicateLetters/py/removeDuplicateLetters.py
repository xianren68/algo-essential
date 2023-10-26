# 定义一个函数，用于移除字符串中的重复字母
class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        # 定义一个栈，用于存储最终结果
        stack = []
        # 定义一个列表，用于存储每个字母出现的次数
        quantity = [0 for _ in range(26)]
        # 定义一个列表，用于存储每个字母是否被使用
        used = [False for _ in range(26)]
        # 遍历字符串中的每一个字母
        for i in s:
            # 记录每个字母出现的次数
            quantity[ord(i) - ord('a')] += 1
        # 遍历字符串中的每一个字母
        for i in s:
            # 如果该字母已被使用，则跳过
            if used[ord(i) - ord('a')]:
                quantity[ord(i) - ord('a')] -= 1
                continue
            # 如果栈不为空，且栈顶字母大于当前字母，且栈顶字母出现的次数大于0，则弹出栈顶字母
            while stack and i < stack[-1] and quantity[ord(stack[-1]) - ord('a')]:
                used[ord(stack[-1]) - ord('a')] = False
                stack.pop()
            # 将当前字母压入栈中
            stack.append(i)
            # 标记当前字母已被使用
            used[ord(i) - ord('a')] = True
            # 记录当前字母出现的次数
            quantity[ord(i) - ord('a')] -= 1
        # 返回最终结果
        return ''.join(stack)