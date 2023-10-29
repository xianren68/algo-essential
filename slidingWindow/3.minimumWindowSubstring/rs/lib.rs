pub fn min_window(s: String, t: String) -> String {
    // 如果s的长度小于t的长度，则返回空字符串
    if s.len() < t.len() {
      return String::from("")
    }
    // 初始化一个长度为256的数组，用于记录t中每个字符的个数
    let mut cnts = [0;256];
    // 遍历t中的每个字符，将其对应的个数减1
    for i in t.chars(){
      cnts[i as usize] -= 1;
    }
    // 初始化左右指针，以及最小窗口大小
    let mut left = 0;
    let mut right = 0;
    let mut all = t.len();
    let mut start:usize = 0;
    let mut min_size = usize::MAX;
    // 将s转换为字符数组
    let chars:Vec<char> = s.chars().collect();
    // 遍历s中的每个字符
    while right < s.len(){
      // 获取当前字符的ASCII码
      let val = chars[right] as usize;
      // 如果t中该字符的个数小于0，则表示该字符在t中出现的次数大于s中出现的次数
      if cnts[val] < 0 {
        all -= 1;
      }
      // 将当前字符的个数加1
      cnts[val]+=1;
      // 如果窗口中字符个数等于t中字符个数，则表示窗口中包含t中所有字符
      if all == 0 {
        // 左指针右移，直到窗口中字符个数刚好满足t中字符个数
        while (cnts[chars[left] as usize]) > 0 {
          cnts[chars[left] as usize] -= 1;
          left += 1
        }
        // 如果当前窗口大小小于最小窗口大小，则更新最小窗口大小和起始位置
        if right-left + 1 < min_size {
          min_size = right-left + 1;
          start = left;
        }
        
      }
      // 右指针右移
      right += 1;
    }
    
    // 如果最小窗口大小等于最大值，则表示不存在符合条件窗口，返回空字符串
    if min_size == usize::MAX {String::from("")} else {String::from(&s[start..start+min_size])}
    
  }