from typing import List


class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        # 栈
        stack = []
        # 矩阵宽度
        n = len(matrix[0])
        # 最大面积
        max_area = 0
        # 高度
        heights = [0] * n
        for val in matrix:
            # 压缩数组
            for i in range(n):
                heights[i] = heights[i] + 1 if val[i] == "1" else 0
                # 如果栈不为空，且当前高度大于栈顶高度，则弹出栈顶元素，计算面积，并更新最大面积
                while stack and heights[stack[-1]] >= heights[i]:
                    top = stack.pop()
                    left = stack[-1] if stack else -1
                    area = heights[top] * (i - left - 1)
                    max_area = max(max_area, area)
                stack.append(i)
        # 清理剩余值
        while stack:
            top = stack.pop()
            left = stack[-1] if stack else -1
            area = heights[top] * (n - left - 1)
            max_area = max(max_area, area)
        return max_area