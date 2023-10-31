from typing import List


class Solution:
    def balancedString(self, s: str) -> int:
        # 计算字符串s的长度
        n = len(s)
        # 创建一个长度为4的列表，用于存储每个字符的个数
        cnts = [0] * 4
        # 创建一个长度为n的列表，用于存储每个字符的索引
        arr = [0] * n
        # 遍历字符串s中的每一个字符
        for i in range(n):
            # 根据字符的不同，分别进行处理
            match s[i]:
                case 'Q':
                    # 如果字符为Q，则将cnts中对应位置的值加1
                    cnts[0] += 1
                    # 将arr中对应位置的值设置为0
                    arr[i] = 0

                case 'W':
                    # 如果字符为W，则将cnts中对应位置的值加1
                    cnts[1] += 1
                    # 将arr中对应位置的值设置为1
                    arr[i] = 1

                case 'E':
                    # 如果字符为E，则将cnts中对应位置的值加1
                    cnts[2] += 1
                    # 将arr中对应位置的值设置为2
                    arr[i] = 2

                case 'R':
                    # 如果字符为R，则将cnts中对应位置的值加1
                    cnts[3] += 1
                    # 将arr中对应位置的值设置为3
                    arr[i] = 3
        # 初始化ans为n
        ans = n
        # 计算target的值
        target = n // 4
        # 初始化l和r为0
        l,r = 0,0
        # 判断cnts中每个元素是否大于target
        if self.ok(cnts, 0, target):
            # 如果大于target，则返回0
            return 0
        # 遍历字符串s中的每一个字符
        for r in range(n):
            # 将cnts中对应位置的值减1
            cnts[arr[r]] -= 1
            # 判断cnts中每个元素是否大于target
            while self.ok(cnts, r-l+1, target):
                # 如果大于target，则将ans的值更新为r-l+1
                ans = min(ans, r-l+1)
                # 将cnts中对应位置的值加1
                cnts[arr[l]] += 1
                # 将l加1
                l += 1
        # 返回ans的值
        return ans
        

    def ok(self, cnts: List[int], leng: int, target: int) -> bool:
        # 遍历cnts中的每一个元素
        for i in range(4):
            # 如果cnts中对应位置的值大于target，则返回False
            if cnts[i] > target:
                return False
            # 判断leng长度能否补齐外面字符
            leng -= target - cnts[i]
        # 如果leng的值等于0，则返回True
        return leng == 0