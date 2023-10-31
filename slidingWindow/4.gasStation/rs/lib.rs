pub fn can_complete_circuit(gas: Vec<i32>, cost: Vec<i32>) -> i32 {
    // 定义左右指针，以及当前累加和
    let mut leng = 0;
    let mut left = 0;
    let mut right = 0;
    let mut sum = 0;
    // 获取数组长度
    let n = cost.len();
    // 循环遍历数组
    while left < n{
      // 重置累加和
      sum = 0;
      // 重置长度
      leng = 0;
      // 循环遍历累加和
      while sum >= 0 {
        // 如果累加和为负数，则返回左指针
        if leng == n{
          return left as i32;
        }
        // 计算右指针
        right = (left + leng)%n;
        // 长度加1
        leng += 1;
        // 累加和加上当前左右指针的差值
        sum += gas[right] - cost[right];
      }
      // 左指针右移
      left += leng;
  }
    // 如果没有找到，则返回-1
    -1
  }