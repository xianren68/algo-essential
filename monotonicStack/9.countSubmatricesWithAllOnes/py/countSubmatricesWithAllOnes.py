from typing import List


class Solution:
    def numSubmat(self, mat: List[List[int]]) -> int:
        # 初始化栈和高度数组
        stack = []
        n = len(mat[0])
        heights = [0] * n
        ans = 0
        # 遍历每一行
        for m in mat:
            # 遍历每一列
            for i in range(n):
                # 更新高度数组
                heights[i] = 0 if m[i] == 0 else heights[i] + 1
                # 如果栈不为空，且当前高度小于栈顶高度，则出栈
                while stack and heights[i] <= heights[stack[-1]]:
                    cur = stack.pop()
                    left = stack[-1] if stack else -1
                    right = i
                    le = right - left - 1
                    # 计算子矩阵的面积
                    quantity = (heights[cur]-max(heights[left] if stack else 0, heights[right])) * le *(le +1)//2
                    # 累加面积
                    ans += quantity
                # 将当前高度入栈
                stack.append(i)
            # 如果栈不为空，则出栈
            while stack:
                    cur = stack.pop()
                    left = stack[-1] if stack else -1
                    right = n
                    le = right - left - 1
                    # 计算子矩阵的面积
                    quantity = (heights[cur]- (heights[left] if stack else 0)) * le *(le +1)//2
                    # 累加面积
                    ans += quantity
        return ans