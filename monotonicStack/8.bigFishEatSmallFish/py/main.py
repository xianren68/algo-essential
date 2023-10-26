# 定义一个函数，参数为n和nums，n为int类型，nums为列表，列表中的元素为int类型
n = int(input())
nums = [int(x) for x in input().split()]
# 定义一个空列表，用于存储栈中的元素
stack = []
# 定义一个变量ans，用于存储最大值
ans = 0
# 遍历nums列表，从最后一个元素开始，每次减1
for i in range(n-1,-1,-1):
    # 定义一个变量cur，用于存储当前元素的最大值
    cur = 0
    # 如果栈不为空，且当前元素大于栈顶元素，则将栈顶元素弹出，并更新cur的值
    while stack and nums[i] > nums[stack[-1]][0]:
        cur = max(cur+1, stack.pop())
    # 将当前元素和cur的值添加到栈中
    stack.append([nums[i],cur])
    # 更新ans的值
    ans = max(cur, ans)
# 输出ans的值
print(ans)