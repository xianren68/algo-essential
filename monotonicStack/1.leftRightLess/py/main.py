# python3会超时，用pypy3测试
# 读取要输入的个数
n = int(input())
nums = list(map(int, input().split()))
# 单调栈
stack = []
# 填充返回值
res = [[]] * n
for i in range(n):
    # 大于等于当前值全出栈
    while len(stack) > 0 and nums[stack[-1]] >= nums[i]:
        # 出栈
        v = stack.pop()
        # 前面不存在比它小的值
        if len(stack) == 0:
            res[v] = [-1,i]
        else:
            res[v] = [stack[-1],i]
    stack.append(i)
# 清理栈中剩余值
while len(stack) > 1:
    v = stack.pop()
    res[v] = [stack[-1],-1]
# 整个数组中最小值
res[stack[0]] = [-1,-1]
# 整理
for i in range(n-2,-1,-1):
    # 处理相同值的情况
    if res[i][1] != -1 and nums[res[i][1]] == nums[i]:
        # 将当前值的右小值变为相同值的右小值
        res[i][1] = res[res[i][1]][1]
# 打印
for i in res:
    print("%d %d" % (i[0],i[1]))