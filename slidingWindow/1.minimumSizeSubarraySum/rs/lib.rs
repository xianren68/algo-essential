impl Solution {
  pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
    // 定义左右指针
    let mut left = 0;
    let mut right = 0;
    // 获取数组的长度
    let length = nums.len() as i32;
    // 初始化结果
    // 设置为比数组长度+1，便于判断整个数组是否有 > target的子数组
    let mut res = length+1;
    // 初始化和
    let mut sum = 0;
    // 右指针小于数组长度时，循环
    while right < length  {
      // 累加和
      sum += nums[right as usize];
      // 当和大于等于目标值时，左指针右移，更新结果
      while sum >= target {
        sum -= nums[left as usize];
        res = min(res,right-left+1);
        left += 1;
      }
      // 右指针右移
      right += 1;
    }
    // 如果结果没有变化，说明没有符合条件的子数组，返回0
    if res == length+1{
      return 0
    }
    // 返回结果
    res
  }
  }
  // 放在impl块外面
  // 定义一个函数，比较两个数的大小
  pub fn min(a:i32,b:i32)->i32{
    // 如果a小于b，返回a
    if a < b {
      return a
    }
    // 否则返回b
    b
  }