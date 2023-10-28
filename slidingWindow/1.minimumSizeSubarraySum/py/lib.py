# 定义一个函数，计算最小子数组的长度
class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        # 定义左右指针
        left, right = 0, 0
        # 计算数组的长度
        l = len(nums)
        # 初始化和
        sum = 0
        # 初始化结果
        res = l + 1
        # 当右指针小于数组长度时，循环
        while right < l:
            # 右指针右移，将右指针指向的元素加入和
            sum += nums[right]
            # 当和大于等于目标值时，更新结果
            while sum >= target:
                res = min(res, right - left + 1)
                # 将左指针指向的元素从和中移除
                sum -= nums[left]
                # 左指针右移
                left += 1
            # 右指针右移
            right += 1
        # 如果结果仍为初始值，说明没有符合条件的子数组，返回0
        if res == l + 1:
            return 0
        # 返回结果
        return res