from typing import List


class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        # 计算gas和cost的长度
        n = len(gas)
        # 定义左右指针
        l,r = 0,0
        # 循环遍历
        while l < n:
            # 定义总油量，以及当前循环的油量
            sum,leng = 0,0
            # 当总油量大于等于0时，循环执行
            while sum >= 0:
                # 如果当前循环的油量等于n，则返回l
                if leng == n:
                    return l
                # 计算右指针
                r = (l + leng) % n
                # 油量加1
                leng += 1
                # 计算总油量
                sum += gas[r] - cost[r]
            # 左指针右移
            l += leng
        # 如果没有满足条件的，则返回-1
        return -1