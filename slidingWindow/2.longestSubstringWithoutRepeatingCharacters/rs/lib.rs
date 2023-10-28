pub fn length_of_longest_substring(s: String) -> i32 {
    // 如果字符串长度为0，则返回0
    if s.len() == 0 {
      return 0
    }
       // 初始化记录数组，记录每个字符最后一次出现的索引
       let mut record = [-1;255];
       // 初始化左右指针
       let mut left = 0;
       let mut right = 0;
       let mut res = 1;
       // 将字符串转换为字符数组
       let chars:Vec<char> = s.chars().collect();
       // 遍历字符数组
       while (right as usize) < s.len(){
         // 获取当前字符的ASCII码值
         let val = chars[right as usize] as usize;
         // 更新左指针，使其指向record数组中val值最后一次出现的索引，如果没有，则指向-1
         left = max(left,record[val] + 1);
         // 将当前字符的索引记录到record数组中
         record[val] = right;
         // 更新结果，使其为当前窗口长度
         res = max(res,right-left+1);
         // 右指针右移
         right += 1
       }
       // 返回结果
       res
  }
  
  // 比较两个数的大小，返回较大的数
  pub fn max(a:i32,b:i32)->i32{
    // 如果a大于b，则返回a，否则返回b
    if a > b {a}else{b}
  }