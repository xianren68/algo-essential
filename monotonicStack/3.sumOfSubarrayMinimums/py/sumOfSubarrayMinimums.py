from typing import List


class Solution:
    MOD = 10 ** 9 + 7
    def sumSubarrayMins(self, arr: List[int]) -> int:
        stack = []
        res = 0
        n = len(arr)
        # 将值全部加入单调栈
        for i in range(n):
            # 如果栈顶大于等于当前值，出栈
            while stack and arr[i] <= arr[stack[-1]]:
                v = stack.pop()
                left = stack[-1] if stack else -1
                res = (res + arr[v] * (v - left) * (i - v)) % self.MOD
            stack.append(i)
        # 清理栈
        while stack:
            v = stack.pop()
            left = stack[-1] if stack else -1
            res = (res + arr[v] * (v - left) * (n - v)) % self.MOD
        return res
        