from typing import List


class Solution:
    def maxWidthRamp(self, nums: List[int]) -> int:
        # 初始化一个空列表，用于存储单调递减的元素下标
        stack = []
        # 获取数组的长度
        n = len(nums)
        # 初始化一个变量，用于存储最大宽度
        ans = 0
        # 遍历数组中的每一个元素
        for i in range(n):
            # 如果栈不为空，且栈顶元素大于等于当前元素，则跳过当前元素
            if stack and nums[stack[-1]] <= nums[i]:
                continue
            # 将当前元素压入栈中
            stack.append(i)
        # 获取数组最后一个元素的索引
        i = n-1
        while stack and i >= 0:
        # 如果栈不为空，且栈顶元素大于等于当前元素，则将栈顶元素弹出，并更新最大宽度
            while stack and nums[stack[-1]] <= nums[i]:
                ans = max(ans, i - stack.pop())
                i -= 1
        # 返回最大宽度
        return ans