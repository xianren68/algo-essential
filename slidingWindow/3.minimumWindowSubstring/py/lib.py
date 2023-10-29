# 定义一个函数，输入参数为字符串s和字符串t，返回值为字符串
class Solution:
    def minWindow(self, s: str, t: str) -> str:
        # 如果s的长度小于t的长度，则返回空字符串
        if len(s) < len(t): return ""
        # 定义左右指针，分别指向s和t的第一个字符
        left,right = 0,0
        # 定义最小长度，初始值为无穷大
        min_len = float("inf")
        # 定义最小字符串的起始位置，初始值为0
        min_start = 0
        # 定义t中字符的总数
        all = len(t)
        # 定义一个数组，用于存储每个字符的个数
        cnts = [0] * 256
        # 遍历t中的每一个字符，将其个数减1
        for c in t:
            cnts[ord(c)] -= 1
        # 遍历s中的每一个字符
        for right in range(len(s)):
            # 如果当前字符在t中，则将其个数加1
            if cnts[ord(s[right])] < 0:
                all -= 1
            cnts[ord(s[right])] += 1
            # 如果t中所有字符的个数都为0，则表示当前窗口符合要求
            if all == 0:
                # 左指针右移，直到找到一个满足条件的窗口
                while cnts[ord(s[left])] > 0:
                    cnts[ord(s[left])] -= 1
                    left += 1
                # 如果当前窗口的长度小于最小窗口的长度，则更新最小窗口的起始位置和长度
                if min_len > right - left + 1:
                    min_len = right - left + 1
                    min_start = left
        # 如果最小窗口的长度未更新，则表示不存在满足条件的窗口，返回空字符串
        if min_len == float("inf"): return ""
        # 返回最小窗口
        return s[min_start:min_start + min_len]