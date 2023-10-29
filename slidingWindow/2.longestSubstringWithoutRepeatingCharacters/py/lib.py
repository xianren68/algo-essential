class Solution:
    def lengthOfLongestSubstring(self,s:str)->int:
        # 如果字符串为空，返回0
        if not s:
            return 0
        # 初始化结果变量res，左右指针left，right，以及一个长度为255的记录数组record
        res = 1
        left,right = 0,0
        record = [-1]*255
        # 右指针小于字符串长度时，循环执行以下操作
        while right < len(s):
            # 左指针left更新为max（left，record[ord(s[right])]+1）
            left = max(left,record[ord(s[right])]+1)
            # res更新为max（res，right-left+1）
            res = max(res,right-left+1)
            # record[ord(s[right])]更新为right
            record[ord(s[right])] = right
            # 右指针右移
            right += 1
        # 返回res
        return res