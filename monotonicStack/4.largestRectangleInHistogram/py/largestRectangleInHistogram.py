from typing import List


class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        max_area = 0
        stack = []
        n = len(heights)
        for i in range(n):
            while stack and heights[stack[-1]] >= heights[i]:
                top = stack.pop()
                left = stack[-1] if stack else -1
                width = i - left - 1    
                area = heights[top] * width
                max_area = max(max_area, area)
            stack.append(i)
        while stack:
            top = stack.pop()
            left = stack[-1] if stack else -1
            width = n - left - 1    
            area = heights[top] * width
            max_area = max(max_area, area)
        return max_area